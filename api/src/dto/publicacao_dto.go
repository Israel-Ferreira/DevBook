package dto

import "errors"

type PublicacaoDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *PublicacaoDTO) Preparar() error {

	if len(p.Title) == 0 && p.Title == "" {
		return errors.New("erro: o titulo não pode estar vazio")
	}

	return nil
}
