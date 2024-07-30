package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"

	_ "github.com/jackc/pgx/v4"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func CreateDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=idp user=postgres password=raptor3796")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
