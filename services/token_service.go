package services

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString([]byte("secret_key"))

	return "dummy_token_for_" + token, err
}
