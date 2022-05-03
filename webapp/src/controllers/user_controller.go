package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

	url := "http://localhost:8990/usuarios"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	responses.JSON(rw,resp.StatusCode, nil)
}
