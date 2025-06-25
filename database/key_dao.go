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
			created_at
		) VALUES ($1, $2, $3)`,
		sigingKey.Kid, sigingKey.PrivateKey, sigingKey.CreatedAt)

	if err != nil {
		slog.Error(fmt.Sprintf("Error saving signing key: %v", err))
	} else {
		sigingKey.Kid = newKid
	}

	return err
}
