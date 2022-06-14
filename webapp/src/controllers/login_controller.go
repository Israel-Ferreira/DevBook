package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/models"
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

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarErro(rw, response)
		return
	}

	var userAuthData models.UserAuth

	if err := json.NewDecoder(response.Body).Decode(&userAuthData); err != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Erro{Erro: err.Error()})
		return
	}

	responses.JSON(rw, response.StatusCode, userAuthData)
}
