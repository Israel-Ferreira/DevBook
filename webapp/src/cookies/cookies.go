package cookies

import (
	"net/http"
	"strconv"

	"github.com/Israel-Ferreira/webapp-devbook/src/config"
	"github.com/gorilla/securecookie"
)

var Cookie *securecookie.SecureCookie

func ConfigurarCookie() {
	Cookie = securecookie.New(config.HashKey, config.BlockKey)
}

func SalvarCookie(rw http.ResponseWriter, ID uint, token string) error {

	id := strconv.Itoa(int(ID))

	dados := map[string]string{
		"id":    id,
		"token": token,
	}

	dadosCode, err := Cookie.Encode("dados", dados)

	if err != nil {
		return err
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "dados",
		Value:    dadosCode,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
