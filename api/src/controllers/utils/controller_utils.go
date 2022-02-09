package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPathIntVar(r *http.Request, pathVariableName string) (int, error) {
	pathVar, err := GetPathVar(r, pathVariableName)

	if err != nil {
		return 0, err
	}

	pathIntVar, err := strconv.ParseInt(pathVar, 10, 64)


	if err != nil {
		return 0, err
	}


	return int(pathIntVar), nil
}

func GetPathVar(r *http.Request, pathVariableName string) (string, error) {
	params := mux.Vars(r)

	pathVariable := params[pathVariableName]

	if pathVariable == "" {
		return "", errors.New("path variable not found")
	}

	return pathVariable, nil
}
