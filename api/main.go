package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/routers"
	"github.com/urfave/negroni/v2"
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

	n := negroni.New(
		negroni.NewLogger(),
	)

	n.UseHandler(r)

	fmt.Printf("Escutando na porta %d \n", config.Porta)

	httpServer := &http.Server{
		Handler: n,
		Addr:    fmt.Sprintf(":%d", config.Porta),
	}

	log.Fatal(httpServer.ListenAndServe())

}
