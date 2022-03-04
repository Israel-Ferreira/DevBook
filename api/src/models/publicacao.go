package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("erro: o titulo não pode estar vazio")
	}

	if p.Conteudo == "" {
		return errors.New("erro: o conteudo do post não deve estar vazio")
	}

	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}

func (p *Publicacao) Preparar() error {

	if erro := p.validar(); erro != nil {
		return erro
	}

	p.formatar()

	return nil
}
