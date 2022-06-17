package middlewares

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/cookies"
)

func Logger(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s /n", r.Method, r.RequestURI, r.Host)
		hf(w, r)
	}
}

func Autenticar(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := cookies.LerCookies(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		hf(w, r)
	}
}
