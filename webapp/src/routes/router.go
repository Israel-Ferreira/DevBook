package routes

import (
	"github.com/Israel-Ferreira/webapp-devbook/src/routes/rotas"
	"github.com/gorilla/mux"
)

func GerarRotas() *mux.Router {
	r := mux.NewRouter()

	r = rotas.ConfigureRoutes(r)

	return r
}
