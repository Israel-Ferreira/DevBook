package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/webapp-devbook/src/routes"
)

func main() {
	fmt.Println("Rodando WebApp")

	router := routes.GerarRotas()

	log.Fatalln(http.ListenAndServe(":3000", router))
}
