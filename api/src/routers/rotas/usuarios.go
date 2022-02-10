package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Uri: "/usuarios", 
		Method: http.MethodGet, 
		RequireAuth: true,
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
		RequireAuth: true,
		Action: controllers.BuscarUsuario,
	},

	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodPut,
		RequireAuth: true,
		Action: controllers.AtualizarUsuario,
	},

	{
		Uri: "/usuarios/{usuarioId}",
		Method: http.MethodDelete,
		RequireAuth: true,
		Action: controllers.DeletarUsuario,
	},

}