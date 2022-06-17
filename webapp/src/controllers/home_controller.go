package controllers

import (
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/utils"
)

func CarregarPaginaPrincipal(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "home.html", nil)
}
