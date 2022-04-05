package controllers

import (
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/utils"
)

func LoadLoginPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "login.html", nil)
}

func LoadCreateUserPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "cadastro.html", nil)
}
