package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri          string
	Method       string
	Action       func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := LoginRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Action).Methods(route.Method)
	}

	return r
}
