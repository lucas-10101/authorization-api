package database

import (
	"authorization-api/models"
	"authorization-api/utils"
	"context"
	"fmt"
	"log/slog"
)

type KeyDataAccessObject struct {
	Connection Connection
	Context    context.Context
}

func (dao *KeyDataAccessObject) SaveKey(sigingKey *models.SigningKey) (err error) {

	var newKid string = utils.GenerateUUID()
	_, err = dao.Connection.ExecContext(dao.Context,
		`INSERT INTO signing_keys (
			kid, 
			private_key, 
			created_at,
			active,
			key_group
		) VALUES ($1, $2, $3, $4, $5)`,
		sigingKey.Kid, sigingKey.PrivateKey, sigingKey.CreatedAt, sigingKey.Active, sigingKey.KeyGroup)

	if err != nil {
		slog.Error(fmt.Sprintf("Error saving signing key: %v", err))
	} else {
		sigingKey.Kid = newKid
	}

	return err
}

func (dao *KeyDataAccessObject) GetCurrentRSASigningKeyByGroup(keyGroup models.SigningKeyGroup) (sigingKey *models.SigningKey, err error) {

	row := dao.Connection.QueryRowContext(dao.Context,
		`SELECT
			kid, 
			private_key, 
			created_at
		FROM 
			signing_keys
		WHERE 
			key_group = $1 AND active = true
		ORDER BY 
			created_at DESC`, keyGroup)

	sigingKey = &models.SigningKey{
		Active:   true,
		KeyGroup: keyGroup,
	}

	err = row.Scan(
		&sigingKey.Kid,
		&sigingKey.PrivateKey,
		&sigingKey.CreatedAt,
	)

	return sigingKey, err
}
