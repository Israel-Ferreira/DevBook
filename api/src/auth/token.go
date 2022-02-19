package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)

	if tokenString == "" {
		return errors.New("empty token")
	}

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if err != nil {
		return err
	}

	fmt.Println(token)

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metódo de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func GenerateSecretKey() (string, error) {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		return "", err
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	return stringBase64, nil
}
