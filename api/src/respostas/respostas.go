package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(rw http.ResponseWriter, status int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		log.Fatalln(err)
	}
}

func Erro(rw http.ResponseWriter, status int, erro error) {
	Json(rw, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
