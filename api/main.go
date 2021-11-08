package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/routers"
)

func main() {
	fmt.Println("Hello Everyone !!!")

	r := routers.GerarRouter()

	routers.LoadHelloRoutes(r)

	httpServer := &http.Server{
		Handler: r,
		Addr:    ":8090",
	}

	log.Fatal(httpServer.ListenAndServe())

}
