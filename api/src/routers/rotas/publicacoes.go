package rotas

import (
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/controllers"
)

var PublishRoutes = []Rota{
	{
		Uri:         "/publicacoes",
		Method:      http.MethodPost,
		Action:      controllers.CriarPublicacao,
		RequireAuth: true,
	},

	{
		Uri:         "/publicacoes",
		Method:      http.MethodGet,
		Action:      controllers.BuscarPublicacoes,
		RequireAuth: true,
	},

	{
		Uri:         "/publicacoes/{publicacaoId}",
		Method:      http.MethodGet,
		Action:      controllers.BuscarPublicacao,
		RequireAuth: true,
	},

	{
		Uri:         "/publicacoes/{publicacaoId}",
		Method:      http.MethodPut,
		Action:      controllers.AtualizarPublicacao,
		RequireAuth: true,
	},

	{
		Uri:         "/publicacoes/{publicacaoId}",
		Method:      http.MethodDelete,
		Action:      controllers.DeletarPublicacao,
		RequireAuth: true,
	},

	{
		Uri:         "/publicacoes/{publicacaoId}/like",
		Method:      http.MethodPut,
		Action:      controllers.CurtirPublicacao,
		RequireAuth: true,
	},
}
