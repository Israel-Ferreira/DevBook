package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
	"github.com/Israel-Ferreira/api-devbook/src/security"
)

func LoginUser(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, err)
		return
	}

	defer r.Body.Close()

	var usuario dto.LoginDTO

	if err = json.Unmarshal(body, &usuario); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	fmt.Println(usuario)

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	user, err := repo.FindByEmail(usuario.Email)

	fmt.Println(err)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	if user.ID == 0 {
		respostas.Erro(rw, http.StatusUnauthorized, errors.New("usuário inválido"))
		return
	}

	fmt.Println(user)

	if err = security.VerificarSenha(user.Senha, usuario.Password); err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	respostas.Json(rw, http.StatusCreated, nil)

}
