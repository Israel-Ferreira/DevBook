package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadHelloRoutes(r *mux.Router) {
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		msg := map[string]string{
			"msg": "Hello World",
		}

		jsonMsg, err := json.Marshal(msg)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Write(jsonMsg)
	}).Methods("GET")
}
