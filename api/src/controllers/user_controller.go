package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Israel-Ferreira/api-devbook/src/auth"
	"github.com/Israel-Ferreira/api-devbook/src/config"
	"github.com/Israel-Ferreira/api-devbook/src/controllers/utils"
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

	if err = usuario.ValidateAndPrepare(http.MethodPost); err != nil {
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
	params := mux.Vars(r)

	usuarioId, err := strconv.ParseInt(params["usuarioId"], 10, 64)

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	if tokenUserId != uint64(usuarioId) {
		respostas.Erro(rw, http.StatusForbidden, errors.New("acesso não autorizado: não é permitido atualizar ou deletar nenhuma conta que não seja sua"))
		return
	}

	corpoReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	var usuario models.Usuario

	if err = json.Unmarshal(corpoReq, &usuario); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	if err = usuario.ValidateAndPrepare(http.MethodPut); err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	if err = repo.AtualizarUsuario(int(usuarioId), usuario); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	usuarioId, err := utils.GetPathIntVar(r, "usuarioId")

	if err != nil {
		respostas.Erro(rw, http.StatusBadRequest, err)
		return
	}

	userId, err := auth.ExtrairUsuarioId(r)

	if err != nil {
		respostas.Erro(rw, http.StatusUnauthorized, err)
		return
	}

	if userId != uint64(usuarioId) {
		respostas.Erro(rw, http.StatusForbidden, errors.New("acesso não autorizado: não é permitido atualizar ou deletar nenhuma conta que não seja sua"))
		return
	}

	db, err := openControllerConnection()

	if err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.UserRepo{Db: db}

	if err = repo.DeletarUsuario(usuarioId); err != nil {
		respostas.Erro(rw, http.StatusInternalServerError, err)
		return
	}

	respostas.Json(rw, http.StatusNoContent, nil)
}

