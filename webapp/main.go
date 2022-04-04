package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/routes"
	"github.com/Israel-Ferreira/webapp-devbook/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp")

	utils.LoadTemplates()

	router := routes.GerarRotas()

	log.Fatalln(http.ListenAndServe(":3000", router))
}
