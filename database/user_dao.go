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
		SELECT id, email, password FROM users WHERE username = $1 AND tenant_id = $3`
	row := dao.Connection.QueryRowContext(dao.Context, query, username, password, tenantId)

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
