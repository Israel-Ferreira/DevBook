package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/controllers"
)

var RotaHome Route = Route{
	Uri:          "/home",
	Method:       http.MethodGet,
	Action:       controllers.CarregarPaginaPrincipal,
	AuthRequired: true,
}
