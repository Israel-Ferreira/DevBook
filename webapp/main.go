package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/config"
	"github.com/Israel-Ferreira/webapp-devbook/src/cookies"
	"github.com/Israel-Ferreira/webapp-devbook/src/routes"
	"github.com/Israel-Ferreira/webapp-devbook/src/utils"
)

func main() {
	config.LoadEnv()

	fmt.Println("Rodando WebApp")
	cookies.ConfigurarCookie()

	utils.LoadTemplates()

	port := fmt.Sprintf(":%d", config.Porta)

	router := routes.GerarRotas()

	log.Println("Servidor iniciando na porta " + port)

	log.Fatalln(http.ListenAndServe(port, router))
}
