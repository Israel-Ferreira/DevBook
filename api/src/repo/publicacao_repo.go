package repo

import (
	"database/sql"

	"github.com/Israel-Ferreira/api-devbook/src/models"
)

type PublicacaoRepo struct {
	db *sql.DB
}

func (pbr PublicacaoRepo) CriarPublicacao(publicacao models.Publicacao) (uint64, error) {
	stmt, err := pbr.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func NovoRepositorioPublicacoes(db *sql.DB) *PublicacaoRepo {
	return &PublicacaoRepo{db}
}
