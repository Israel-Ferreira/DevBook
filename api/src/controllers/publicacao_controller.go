package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/models"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
)

func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {}

func BuscarPublicacoes(rw http.ResponseWriter, r *http.Request) {}

func DeletarPublicacao(rw http.ResponseWriter, r *http.Request) {}

func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {}

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

	publicacao.ID, err = repo.CriarPublicacao(publicacao)

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	respostas.Json(rw, http.StatusCreated, publicacao)

}
