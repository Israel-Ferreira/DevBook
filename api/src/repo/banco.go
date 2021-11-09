package repo

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenConnection(connStr string) {
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatalln("Erro ao criar a conexão com banco")
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Erro ao tentar conectar com o banco")
	}

	DB = db
}
