package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/controllers/utils"
	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
	"github.com/Israel-Ferreira/api-devbook/src/security"
)

func ChangeUserPassword(rw http.ResponseWriter, r *http.Request) {
	usuarioIDToken, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	usuarioId, err := utils.GetPathIntVar(r, "usuarioId")

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if usuarioIDToken != uint64(usuarioId) {
		respostas.Erro(rw, http.StatusForbidden, errors.New("erro: não é possivel atualizar um usuário que não seja o seu"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	var senhaReq dto.ChangePasswordDTO

	if err = json.Unmarshal(body, &senhaReq); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NewUserRepo(db)

	password, err := repo.BuscarSenha(uint(usuarioId))

	if err != nil {
		respostas.Erro(rw, http.StatusNotFound, err)
		return
	}

	if err = security.VerificarSenha(password, senhaReq.CurrentPassword); err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	senhaHash, err := security.HashPassword(senhaReq.NewPassword)

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	if err = repo.AtualizarSenha(uint(usuarioId), string(senhaHash)); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)

}
