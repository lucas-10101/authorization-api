package services

import (
	"authorization-api/database"
	"authorization-api/models"
	"authorization-api/utils"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func GenerateNewRSASigingKey() (sigingKey *models.SigningKey, err error) {

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
	}

	var keyDao = &database.KeyDataAccessObject{
		Connection: database.GetDatabaseConnection(),
		Context:    context.TODO(),
	}
	err = keyDao.SaveKey(sigingKey)

	return sigingKey, err
}
