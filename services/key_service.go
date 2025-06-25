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
	"time"
)

func GenerateNewRSASigingKey() (sigingKey *models.SigningKey, err error) {

	var privateKey *rsa.PrivateKey
	if privateKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		slog.Error(fmt.Sprintf("Error generating RSA key: %v", err))
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}

	var privateKeyContent string = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}))

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
