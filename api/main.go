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

	db, err := config.OpenConnection(config.ConexaoDbString)

	if err != nil {
		log.Fatalln("Erro ao abrir a conexão com o banco de dados")
	}

	defer db.Close()

	routers.LoadHelloRoutes(r)

	fmt.Printf("Escutando na porta %d \n", config.Porta)

	httpServer := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", config.Porta),
	}

	log.Fatal(httpServer.ListenAndServe())

}
