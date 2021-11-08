package routers

import (
	"github.com/Israel-Ferreira/api-devbook/src/routers/rotas"
	"github.com/gorilla/mux"
)

func GerarRouter() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configuration(r)
}
