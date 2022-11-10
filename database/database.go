package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver conect
)

func Connect() (*sql.DB, error) {
	// user: golang, pass: golang, database: devbook
	connectDB := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", connectDB)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
