package controllers

import "net/http"

func LoadLoginPage(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Tela de Login"))
}
