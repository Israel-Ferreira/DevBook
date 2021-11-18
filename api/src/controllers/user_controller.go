package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/models"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
	"github.com/gorilla/mux"
)

func openControllerConnection() (*sql.DB, error) {
	db, err := config.OpenConnection(config.ConexaoDbString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {

	db, err := openControllerConnection()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	var usuario models.Usuario

	corpoReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(corpoReq, &usuario); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if err = usuario.ValidateAndPrepare(); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if err = repo.AddUsuario(usuario); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	respostas.Json(rw, 201, usuario)
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(params["usuarioId"], 10, 64)

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	dbRepo := repo.UserRepo{Db: db}

	usuario, err := dbRepo.BuscarUsuarioPorId(int(usuarioID))

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}


	if usuario.ID == 0 && err == nil {
		respostas.Erro(rw, http.StatusNotFound, errors.New("usuario não encontrado"))
		return
	}

	respostas.Json(rw, http.StatusOK, usuario)
}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	usuarios, err := repo.BuscarUsuarios(nomeOuNick)

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusOK, usuarios)
}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Atualizar usuario por id")
	rw.Write([]byte("Atualizar usuario por id "))
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletar usuario por id")
	rw.Write([]byte("Deletar Usuario"))
}
