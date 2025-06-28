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

	permissionDao := database.PermissionDataAccessObject{
		Connection: s.Connection,
		Context:    s.Context,
	}

	var permissions []*models.Permission
	if permissions, err = permissionDao.FindAllByUserID(user.ID); err != nil {
		return "", err
	}

	var scopedRoles []string = make([]string, len(permissions))
	for index, permission := range permissions {
		scopedRoles[index] = permission.ToScopedResource()
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
		"scopes":    scopedRoles,
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

func (s *AuthenticationService) MakeScopeList(permissions []*models.Permission) []string {
	var roles []string
	for _, perm := range permissions {
		roles = append(roles, perm.Scope)
	}
	return roles
}

func (s *AuthenticationService) ValidateToken(token string) (*jwt.Token, error) {

	keyService := &SigningKeyService{
		Connection: s.Connection,
		Context:    s.Context,
	}
	signingKey, err := keyService.GetCurrentRSASigningKeyByGroup(models.JWT_KEY_GROUP)
	if err != nil {
		return nil, err
	}

	tokenData, err := jwt.Parse(
		token,
		func(token *jwt.Token) (any, error) {

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, jwt.ErrTokenInvalidClaims
			}

			requiredClaims := []string{"tenant_id", "scopes", "email"}

			for _, claim := range requiredClaims {
				if _, exists := claims[claim]; !exists {
					return nil, jwt.ErrTokenInvalidClaims
				}
			}

			return &signingKey.RsaPrivateKey.PublicKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}),
		jwt.WithIssuer(os.Getenv("APPLICATION_NAME")),
		jwt.WithExpirationRequired(),
	)

	return tokenData, err
}
