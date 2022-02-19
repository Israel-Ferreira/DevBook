package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

func parseToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extrairToken(r)

	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidarToken(r *http.Request) error {

	token, err := parseToken(r)

	if err != nil {
		return err
	}

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

func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	token, err := parseToken(r)

	if err != nil {
		return 0, err
	}

	if permitions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(permitions["usuarioId"])
		usuarioIdPermitions := fmt.Sprintf("%.0f", permitions["usuarioId"])

		usuarioID, err := strconv.ParseUint(usuarioIdPermitions, 10, 64)

		if err != nil {
			return 0, err
		}

		return usuarioID, nil
	}

	return 0, errors.New("cannot extract token")
}

func GenerateSecretKey() (string, error) {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		return "", err
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	return stringBase64, nil
}
