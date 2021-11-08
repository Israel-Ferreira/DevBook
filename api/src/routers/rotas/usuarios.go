package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Uri: "/usuarios", 
		Method: http.MethodGet, 
		RequireAuth: false,
		Action: controllers.BuscarUsuarios,
	},

	{
		Uri: "/usuarios",
		Method: http.MethodPost,
		RequireAuth: false,
		Action: controllers.CriarUsuario,
	},

	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodGet,
		RequireAuth: false,
		Action: controllers.BuscarUsuario,
	},

	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodPut,
		RequireAuth: false,
		Action: controllers.AtualizarUsuario,
	},

	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodDelete,
		RequireAuth: false,
		Action: controllers.DeletarUsuario,
	},

}