package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/routers"
)

func main() {
	fmt.Println("Hello Everyone !!!")

	r := routers.GerarRouter()


	config.Carregar()

	routers.LoadHelloRoutes(r)

	fmt.Println(config.ConexaoDbString)
	fmt.Println(config.Porta)

	httpServer := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", config.Porta),
	}

	log.Fatal(httpServer.ListenAndServe())

}
