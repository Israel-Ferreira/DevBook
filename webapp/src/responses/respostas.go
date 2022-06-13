package responses

import (
	"encoding/json"
	"net/http"
)

type Erro struct {
	Erro string `json:"erro"`
}

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(int(statusCode))

	if err := json.NewEncoder(rw).Encode(&data); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func TratarErro(rw http.ResponseWriter, r *http.Response) {
	var errorResponse Erro

	if err := json.NewDecoder(r.Body).Decode(&errorResponse); err != nil {
		JSON(rw, http.StatusInternalServerError, Erro{Erro: err.Error()})
		return
	}

	JSON(rw, r.StatusCode, errorResponse)

}
