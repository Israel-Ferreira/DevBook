package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/controllers/utils"
	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/models"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
)

func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {}

func BuscarPublicacoes(rw http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepositorioPublicacoes(db)

	publicacoes, err := repo.BuscarPublicacoes(uint(usuarioID))

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusOK, publicacoes)

}

func DeletarPublicacao(rw http.ResponseWriter, r *http.Request) {}

func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {
	publicacaoId, err := utils.GetPathIntVar(r, "publicacaoId")

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepositorioPublicacoes(db)

	publicacao, err := repo.BuscarPublicacao(uint(publicacaoId))

	if err != nil {
		respostas.Erro(rw, http.StatusNotFound, err)
		return
	}

	respostas.Json(rw, http.StatusOK, publicacao)
}

func CriarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioId, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	var publicacaoDTO dto.PublicacaoDTO

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&publicacaoDTO); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepositorioPublicacoes(db)

	publicacao := models.Publicacao{
		AutorId:  usuarioId,
		Titulo:   publicacaoDTO.Title,
		Conteudo: publicacaoDTO.Content,
	}

	if err := publicacao.Preparar(); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	publicacao.ID, err = repo.CriarPublicacao(publicacao)

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusCreated, publicacao)

}
