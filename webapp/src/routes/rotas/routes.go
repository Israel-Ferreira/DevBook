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

	routes = append(routes, RotaHome)

	routes = append(routes, UserRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Action).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
