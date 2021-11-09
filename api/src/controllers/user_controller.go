package controllers

import (
	"fmt"
	"net/http"
)

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando usuario")
	rw.Write([]byte("Criando usuario"))
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando um usuário na base")
	rw.Write([]byte("Buscando o usuario por id"))
}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando todos usuarios da base")
	rw.Write([]byte("Buscando todos os usuarios "))
}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Atualizar usuario por id")
	rw.Write([]byte("Atualizar usuario por id "))
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletar usuario por id")
	rw.Write([]byte("Deletar Usuario"))
}
