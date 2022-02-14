package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("[DevBook] %s %s \n", r.URL, r.Method)
		next(rw, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")

		if err := auth.ValidarToken(r); err != nil {
			respostas.Erro(rw, http.StatusUnauthorized, err)
			return
		}

		next(rw, r)
	}
}
