package database

import (
	"authorization-api/models"
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserDataAccessObject struct {
	Connection Connection
	Context    context.Context
}

// Find User By Username, Password, and TenantId
func (dao UserDataAccessObject) FindUserByUsernameAndTenantId(username, password, tenantId string) (user *models.User, err error) {

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	tenantId = strings.TrimSpace(tenantId)

	if username == "" || password == "" || tenantId == "" {
		return nil, errors.New("username, password, and tenantId cannot be empty")
	}

	query := `
		SELECT id, email, password FROM users WHERE username = $1 AND tenant_id = $2`
	row := dao.Connection.QueryRowContext(dao.Context, query, username, tenantId)

	user = &models.User{}

	user.Username = username
	user.TenantId = tenantId
	user.Password = "[REDACTED]" // Do not return the password in the user object

	var passwordHash string = ""
	err = row.Scan(
		&user.ID,
		&user.Email,
		&passwordHash,
	)

	if err != nil || passwordHash == "" || bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) != nil {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (dao UserDataAccessObject) FindUserById(userId string) (user *models.User, err error) {
	if userId == "" {
		return nil, errors.New("userId cannot be empty")
	}

	row := dao.Connection.QueryRowContext(
		dao.Context,
		`SELECT
			ID,
			USERNAME,
			EMAIL,
			TENANT_ID
		FROM
			USERS
		WHERE
			ID = $1`,
		userId)

	user = &models.User{}
	user.Password = "[REDACTED]" // Do not return the password in the user object
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.TenantId,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
