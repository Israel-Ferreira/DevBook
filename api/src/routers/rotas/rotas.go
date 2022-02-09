package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri         string
	Method      string
	Action      func(rw http.ResponseWriter, r *http.Request)
	RequireAuth bool
}

func NewRouteWithoutAuth(uri string, method string, action func(rw http.ResponseWriter, r *http.Request)) Rota {
	return Rota{
		Uri:         uri,
		Method:      method,
		Action:      action,
		RequireAuth: false,
	}
}

func Configuration(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	rotas = append(rotas, LoginRota())

	for _, userRoute := range rotas {
		r.HandleFunc(userRoute.Uri, userRoute.Action).Methods(userRoute.Method)
	}

	

	return r
}
