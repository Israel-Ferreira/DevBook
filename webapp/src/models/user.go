package models

type User struct {
	Nome  string `json:"nome"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
