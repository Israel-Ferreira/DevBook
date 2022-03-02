package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Uri:         "/usuarios",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.BuscarUsuarios,
	},

	{
		Uri:         "/usuarios",
		Method:      http.MethodPost,
		RequireAuth: false,
		Action:      controllers.CriarUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.BuscarUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}",
		Method:      http.MethodPut,
		RequireAuth: true,
		Action:      controllers.AtualizarUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}",
		Method:      http.MethodDelete,
		RequireAuth: true,
		Action:      controllers.DeletarUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}/followers",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.SeguirUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}/followers",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.BuscarSeguidores,
	},

	{
		Uri:         "/usuarios/{usuarioId}/unfollow",
		Method:      http.MethodDelete,
		RequireAuth: true,
		Action:      controllers.UnfollowUsuario,
	},

	{
		Uri:         "/usuarios/{usuarioId}/following",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.BuscarSeguindo,
	},

	{
		Uri:         "/usuarios/{usuarioId}/change-password",
		Method:      http.MethodPost,
		Action:      controllers.ChangeUserPassword,
		RequireAuth: true,
	},
}
