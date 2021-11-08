package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/routers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Everyone !!!")

	r := mux.NewRouter()

	routers.LoadHelloRoutes(r)


	httpServer := &http.Server{
		Handler: r,
		Addr: ":8090",
	}

	log.Fatal(httpServer.ListenAndServe())

}
