package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/config"
	"github.com/Israel-Ferreira/webapp-devbook/src/responses"
)

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	jsonForm := map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	}

	usuario, err := json.Marshal(jsonForm)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Erro{Erro: err.Error()})
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode >= 400 {
		errBadReq := errors.New("erro ao validar o usuario")
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Erro{Erro: errBadReq.Error()})
		return
	}

	responses.JSON(rw, resp.StatusCode, nil)
}
