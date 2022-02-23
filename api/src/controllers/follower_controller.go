package controllers

import (
	"errors"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/controllers/utils"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
)

func BuscarSeguidores(rw http.ResponseWriter, r *http.Request) {
	usuarioID, err := utils.GetPathIntVar(r, "usuarioId")

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

	repo := repo.NewUserRepo(db)

	followers, err := repo.BuscarSeguidoresDoUsuario(usuarioID)

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusOK, followers)

}

func UnfollowUsuario(rw http.ResponseWriter, r *http.Request) {
	seguidorID, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	usuarioID, err := utils.GetPathIntVar(r, "usuarioId")

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if usuarioID == int(seguidorID) {
		respostas.Erro(rw, http.StatusForbidden, errors.New("você não pode parar de seguir a si mesmo"))
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NewUserRepo(db)

	if err = repo.PararDeSeguir(usuarioID, int(seguidorID)); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)

}

func SeguirUsuario(rw http.ResponseWriter, r *http.Request) {
	seguidorID, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	usuarioId, err := utils.GetPathIntVar(r, "usuarioId")

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if usuarioId == int(seguidorID) {
		respostas.Erro(rw, http.StatusForbidden, errors.New("você não pode seguir a si mesmo"))
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NewUserRepo(db)

	if err = repo.SeguirUsuario(usuarioId, int(seguidorID)); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)
}
