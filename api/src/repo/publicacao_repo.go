package repo

import (
	"database/sql"

	"github.com/Israel-Ferreira/api-devbook/src/models"
)

type PublicacaoRepo struct {
	db *sql.DB
}

func (pbr PublicacaoRepo) BuscarPublicacao(id uint) (models.Publicacao, error) {
	query, err := pbr.db.Query(`
		select p.*, u.nick from publicacoes 
		p inner join usuarios u
		on u.id = p.autor_id
		where p.id = ?`,
		id,
	)

	if err != nil {
		return models.Publicacao{}, err
	}

	defer query.Close()

	var publicacao models.Publicacao

	if query.Next() {
		if err := query.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorId, &publicacao.Curtidas, &publicacao.CriadoEm, &publicacao.AutorNick); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacao, nil
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
