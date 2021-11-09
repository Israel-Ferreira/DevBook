package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection(connStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
