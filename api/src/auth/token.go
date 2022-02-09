package auth

import (
	"time"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	jwt "github.com/dgrijalva/jwt-go/v4"
)

func CriarToken(usuarioId uint64) (string, error) {
	permitions := jwt.MapClaims{}

	permitions["authorized"] = true
	permitions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permitions["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)

	return token.SignedString(config.SecretKey)
}
