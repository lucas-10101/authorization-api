package services

import (
	"authorization-api/database"
	"authorization-api/models"
	"authorization-api/utils"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type SigningKeyService struct {
	Connection database.Connection
	Context    context.Context
}

// GenerateNewRSASigingKey generates a new RSA signing key, encrypts it, and saves it to the database.
func (s *SigningKeyService) GenerateNewRSASigingKey() (sigingKey *models.SigningKey, err error) {

	var privateKey *rsa.PrivateKey
	var bitSize int

	if bitSize, err = strconv.Atoi(os.Getenv("RSA_PRIVATE_KEY_BIT_SIZE")); err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key bit size: %w", err)
	}

	if privateKey, err = rsa.GenerateKey(rand.Reader, bitSize); err != nil {
		slog.Error(fmt.Sprintf("Error generating RSA key: %v", err))
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}

	var privateKeyContent string = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}))

	if privateKeyContent, err = utils.EncryptAES(privateKeyContent); err != nil {
		slog.Error(fmt.Sprintf("Error encrypting RSA private key: %v", err))
		return nil, fmt.Errorf("failed to encrypt RSA private key: %w", err)
	}

	sigingKey = &models.SigningKey{
		Kid:        utils.GenerateUUID(),
		CreatedAt:  time.Now().Format(time.RFC3339),
		PrivateKey: privateKeyContent,
		Active:     true,
		KeyGroup:   models.JWT_KEY_GROUP,
	}

	var keyDao = &database.KeyDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}
	err = keyDao.SaveKey(sigingKey)

	return sigingKey, err
}

// GetCurrentRSASigningKeyByGroup retrieves the current RSA signing key for a given key group.
// If no active key is found, it generates a new RSA signing key.
// It returns the signing key with the decrypted private key.
func (s *SigningKeyService) GetCurrentRSASigningKeyByGroup(keyGroup models.SigningKeyGroup) (sigingKey *models.SigningKey, err error) {

	if keyGroup == "" {
		return nil, fmt.Errorf("key group cannot be empty")
	}

	var keyDao = &database.KeyDataAccessObject{
		Connection: database.GetDatabaseConnection(),
		Context:    s.Context,
	}

	sigingKey, err = keyDao.GetCurrentRSASigningKeyByGroup(keyGroup)
	if err != nil && err != sql.ErrNoRows {
		slog.Error(fmt.Sprintf("Error getting current RSA signing key: %v", err))
		return nil, fmt.Errorf("failed to get current RSA signing key: %w", err)
	}

	if err == sql.ErrNoRows {
		slog.Info("No active RSA signing key found, generating a new one")
		sigingKey, err = s.GenerateNewRSASigingKey()

		if err != nil {
			slog.Error(fmt.Sprintf("Error generating new RSA signing key: %v", err))
			return nil, fmt.Errorf("failed to generate new RSA signing key: %w", err)
		}
	}

	decryptedKey, err := utils.DecryptAES(sigingKey.PrivateKey)

	if err != nil {
		slog.Error(fmt.Sprintf("Error decrypting AES RSA private key: %v", err))
		return nil, fmt.Errorf("failed to decrypt AES RSA private key: %w", err)
	}

	block, _ := pem.Decode([]byte(decryptedKey))

	if block == nil || block.Type != "RSA PRIVATE KEY" {
		slog.Error("Failed to decode RSA private key PEM block")
		return nil, fmt.Errorf("failed to decode RSA private key PEM block")
	}

	sigingKey.RsaPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)

	return sigingKey, err
}
