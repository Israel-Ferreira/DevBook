package repo

import (
	"database/sql"

	"github.com/Israel-Ferreira/api-devbook/src/models"
)

type UserRepo struct {
	Db *sql.DB
}

func (u UserRepo) GetUsuarios() ([]models.Usuario, error) {
	query, err := u.Db.Query("select id, nome, email, nick from usuarios")

	usuarios := []models.Usuario{}

	if err != nil {
		return nil, err
	}

	defer query.Close()

	for query.Next() {
		var usuario models.Usuario

		if erro := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u UserRepo) AddUsuario(user models.Usuario) error {

	stmt, err := u.Db.Prepare(
		`insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)`,
	)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.Nome,
		user.Nick,
		user.Email,
		user.Senha,
	)

	if err != nil {
		return err
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	return nil
}
