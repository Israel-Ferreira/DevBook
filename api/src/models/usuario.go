package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (u *Usuario) ValidateAndPrepare() error {
	if erro := u.validar(); erro != nil {
		return erro
	}

	u.formatar()

	return nil
}

func (u Usuario) validar() error {
	if u.Nome == "" {
		return errors.New("o campo nome é obrigatorio e não pode estar vazio")
	}

	if u.Nick == "" {
		return errors.New("o campo nick é obrigatorio e não pode ficar vazio")
	}

	if u.Email == "" {
		return errors.New("o campo email é obrigatorio e não pode estar em branco")
	}

	if u.Senha == "" {
		return errors.New("a senha é obrigatoria e não pode estar em branco")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)

	u.Email = strings.TrimSpace(u.Email)
}
