package data

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	db *sql.DB
}

func CreateDatabase() (Database, error) {
	db, err := sql.Open("pgx", "user=postgres password=admin host=localhost port=5432 database=postgres sslmode=disable")
	return Database{db}, err
}
