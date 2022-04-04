package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/controllers/utils"
	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/models"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
	"github.com/gorilla/mux"
)

func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	publicacaoId, err := strconv.ParseInt(parametros["publicacaoId"], 10, 64)

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

	pubRepo := repo.NovoRepositorioPublicacoes(db)

	publicacao, err := pubRepo.BuscarPublicacao(uint(publicacaoId))

	if err != nil {
		respostas.Erro(rw, http.StatusNotFound, err)
		return
	}

	if publicacao.AutorId != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("acesso não autorizado"))
		return
	}

	var body dto.PublicacaoDTO

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	if err = body.Preparar(); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if err = pubRepo.AtualizarPublicacao(uint(publicacaoId), body); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)
}

func BuscarPublicacoesDoUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["usuarioId"])

	if err != nil {
		respostas.Json(rw, http.StatusNotFound, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepositorioPublicacoes(db)

	publicacoes, err := repo.BuscarPublicacoesDoUsuario(uint(userId))

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusOK, publicacoes)
}

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

func DeletarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	publicacaoId, err := strconv.ParseInt(parametros["publicacaoId"], 10, 64)

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

	pubRepo := repo.NovoRepositorioPublicacoes(db)

	publicacao, err := pubRepo.BuscarPublicacao(uint(publicacaoId))

	if err != nil {
		respostas.Erro(rw, http.StatusNotFound, err)
		return
	}

	if publicacao.AutorId != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("acesso não autorizado"))
		return
	}

	if err = pubRepo.DeletarPublicacao(uint(publicacao.ID)); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)

}

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

func CurtirPublicacao(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)

	if err != nil {
		respostas.Json(rw, http.StatusBadRequest, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepositorioPublicacoes(db)

	if err = repo.CurtirPublicacao(uint(publicacaoID)); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)
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
