package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/controllers"
)

var LoginRoutes = []Route{
	{
		Uri:          "/",
		Method:       http.MethodGet,
		Action:       controllers.LoadLoginPage,
		AuthRequired: false,
	},

	{
		Uri:          "/login",
		Method:       http.MethodGet,
		Action:       controllers.LoadLoginPage,
		AuthRequired: false,
	},
}
