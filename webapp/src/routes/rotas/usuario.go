package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/controllers"
)

var UserRoutes = []Route{
	{
		Uri:    "/create-user",
		Method: http.MethodGet,
		Action: controllers.LoadCreateUserPage,
		AuthRequired: false,
	},
}


