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

func LerCookies(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("dados")

	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err := Cookie.Decode("dados", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
