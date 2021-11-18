package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/models"
	"github.com/Israel-Ferreira/api-devbook/src/repo"
	"github.com/Israel-Ferreira/api-devbook/src/respostas"
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

	if err = repo.AddUsuario(usuario); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	respostas.Json(rw, 201, corpoReq)
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando um usuário na base")
	rw.Write([]byte("Buscando o usuario por id"))
}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	usuarios, err := repo.GetUsuarios()

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
