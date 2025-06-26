package services

import (
	"authorization-api/database"
	"authorization-api/models"
	"context"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationService struct {
	Connection database.Connection
	Context    context.Context
}

func (s *AuthenticationService) AuthenticateUser(username, password, tenantId string) (jwtTokenString string, err error) {

	var userDao *database.UserDataAccessObject
	var user *models.User

	userDao = &database.UserDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	if user, err = userDao.FindUserByUsernameAndTenantId(username, password, tenantId); err != nil {
		return "", err
	}

	var tokenTTL int64
	if tokenTTL, err = strconv.ParseInt(os.Getenv("JWT_TOKEN_LIFESPAN_MINUTES"), 10, 64); err != nil {
		tokenTTL = 1
	}

	var tokenModel *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":       os.Getenv("APPLICATION_NAME"),
		"sub":       user.ID,
		"exp":       time.Now().Add(time.Minute * time.Duration(tokenTTL)).Unix(),
		"email":     user.Email,
		"tenant_id": user.TenantId,
	})

	keyService := &SigningKeyService{
		Connection: s.Connection,
		Context:    s.Context,
	}
	var signingKey *models.SigningKey
	if signingKey, err = keyService.GetCurrentRSASigningKeyByGroup(models.JWT_KEY_GROUP); err != nil {
		return "", err
	}

	jwtTokenString, err = tokenModel.SignedString(signingKey.RsaPrivateKey)

	return jwtTokenString, err
}
