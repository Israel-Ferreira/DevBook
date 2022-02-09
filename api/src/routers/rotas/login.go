package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/controllers"
)

func LoginRota() Rota {
	return Rota{
		Uri:         "/login",
		Method:      http.MethodPost,
		Action:      controllers.LoginUser,
		RequireAuth: false,
	}
}
