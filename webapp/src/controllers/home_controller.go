package controllers

import (
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/config"
	"github.com/Israel-Ferreira/webapp-devbook/src/requests"
	"github.com/Israel-Ferreira/webapp-devbook/src/utils"
)

func CarregarPaginaPrincipal(rw http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publicacoes", config.ApiUrl)

	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	fmt.Println(response)

	fmt.Println(erro)

	fmt.Println(url)
	utils.RenderTemplate(rw, "home.html", nil)
}
