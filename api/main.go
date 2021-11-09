package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/routers"
)

func main() {
	fmt.Println("Hello Everyone !!!")

	r := routers.GerarRouter()

	config.Carregar()

	repo.OpenConnection(config.ConexaoDbString)

	defer repo.DB.Close()

	routers.LoadHelloRoutes(r)

	fmt.Printf("Escutando na porta %d \n", config.Porta)

	httpServer := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", config.Porta),
	}

	log.Fatal(httpServer.ListenAndServe())

}
