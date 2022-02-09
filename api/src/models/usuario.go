package models

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Israel-Ferreira/api-devbook/src/security"
	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (u *Usuario) ValidateAndPrepare(httpMethod string) error {

	if erro := u.validar(httpMethod); erro != nil {
		return erro
	}


	if erro := u.formatar(httpMethod); erro != nil {
		return erro
	}


	return nil
}

func (u Usuario) validar(httpMethod string) error {
	if u.Nome == "" {
		return errors.New("o campo nome é obrigatorio e não pode estar vazio")
	}

	if u.Nick == "" {
		return errors.New("o campo nick é obrigatorio e não pode ficar vazio")
	}

	if u.Email == "" {
		return errors.New("o campo email é obrigatorio e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o email inserido é invalido")
	}

	if httpMethod == http.MethodPost && u.Senha == "" {
		return errors.New("a senha é obrigatoria e não pode estar em branco")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)

	u.Email = strings.TrimSpace(u.Email)

	if etapa == http.MethodPost {
		senhaHasheada, err := security.HashPassword(u.Senha)

		if err != nil {
			return err
		}

		u.Senha = string(senhaHasheada)
	}


	return nil
}
