package repo

import (
	"database/sql"
	"fmt"

	"github.com/Israel-Ferreira/api-devbook/src/dto"
	"github.com/Israel-Ferreira/api-devbook/src/models"
)

type UserRepo struct {
	Db *sql.DB
}

func (u UserRepo) FindByEmail(email string) (models.Usuario, error) {
	query, err := u.Db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)

	if err != nil {
		return models.Usuario{}, err
	}

	defer query.Close()

	var usuario models.Usuario

	if query.Next() {
		if err = query.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (u UserRepo) AtualizarUsuario(id int, usuarioAt models.Usuario) error {
	stmt, erro := u.Db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer stmt.Close()

	if _, erro = stmt.Exec(usuarioAt.Nome, usuarioAt.Nick, usuarioAt.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (u UserRepo) BuscarUsuarioPorId(id int) (models.Usuario, error) {
	query, err := u.Db.Query(
		"select id, nome, email, nick, criadoEm from usuarios where id = ?",
		id,
	)

	if err != nil {
		return models.Usuario{}, err
	}

	defer query.Close()

	var usuario models.Usuario

	if query.Next() {
		if err = query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Nick, &usuario.CriadoEm); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (u UserRepo) BuscarUsuarios(username string) ([]models.Usuario, error) {
	nomeOrNick := fmt.Sprintf("%%%s%%", username)

	query, err := u.Db.Query(
		"select id, nome, email, nick, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOrNick,
		nomeOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	var usuarios []models.Usuario

	for query.Next() {
		var usuario models.Usuario
		if err = query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Nick, &usuario.CriadoEm); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
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

func (u UserRepo) DeletarUsuario(id int) error {
	stmt, err := u.Db.Prepare("delete from usuarios where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil

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

func (u UserRepo) PararDeSeguir(usuarioID, seguidorID int) error {
	stmt, err := u.Db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (u UserRepo) SeguirUsuario(usuarioID, seguidorID int) error {
	stmt, err := u.Db.Prepare(
		"insert ignore into seguidores(usuario_id, seguidor_id) values (?, ?)",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) BuscarSeguidoresDoUsuario(usuarioID int) ([]*dto.FollowerDTO, error) {
	followers := []*dto.FollowerDTO{}

	query, err := u.Db.Query(`
		select u.id, u.nome, u.nick, u.email from usuarios u
		inner join seguidores s on s.seguidor_id = u.id
		where s.usuario_id = ?`,
		usuarioID,
	)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	for query.Next() {
		var follower dto.FollowerDTO

		if err = query.Scan(&follower.ID, &follower.Username, &follower.Nick, &follower.Email); err != nil {
			return nil, err
		}

		followers = append(followers, &follower)
	}

	return followers, nil
}

func (u UserRepo) BuscarSeguindo(usuarioID uint) ([]*dto.FollowingUserDTO, error) {
	query, err := u.Db.Query(`
		select u.id, u.nome, u.nick, u.email from usuarios u
		inner join seguidores s on u.id = s.usuario_id
		where s.seguidor_id = ?`,
		usuarioID,
	)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	var usuarios []*dto.FollowingUserDTO

	for query.Next() {
		var usuario dto.FollowingUserDTO
		if err := query.Scan(&usuario.ID, &usuario.Username, &usuario.Nick, &usuario.Email); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, &usuario)
	}

	return usuarios, nil
}

func (u UserRepo) BuscarSenha(usuarioID uint) (string, error) {
	query, err := u.Db.Query("select senha from usuarios where id = ?", usuarioID)

	if err != nil {
		return "", err
	}

	defer query.Close()

	var usuario models.Usuario

	if query.Next() {
		if err := query.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}

	return usuario.Senha, nil
}

func (u UserRepo) AtualizarSenha(usuarioID uint, hashPassword string) error {
	statement, err := u.Db.Prepare("update usuarios set senha = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(hashPassword, usuarioID); err != nil {
		return err
	}

	return nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{Db: db}
}
