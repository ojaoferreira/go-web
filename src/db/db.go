package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Conectar() *sql.DB {
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/alura_loja?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != err {
		panic(err)
	}
	return db
}
