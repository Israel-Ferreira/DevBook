package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/responses"
)

func LoginUser(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Erro{Erro: err.Error()})
		return
	}

	fmt.Println(string(usuario))

	url := "http://localhost:8990/login"

	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Erro{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {

		fmt.Println(response.StatusCode)
		errBadReq := errors.New("erro ao validar o usuario")
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Erro{Erro: errBadReq.Error()})
		return
	}

	token, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(token))

	responses.JSON(rw, response.StatusCode, response.Body)
}
