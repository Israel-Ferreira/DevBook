package responses

import (
	"encoding/json"
	"net/http"
)

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(int(statusCode))

	if err := json.NewEncoder(rw).Encode(&data); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	
}
