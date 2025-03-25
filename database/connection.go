package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
)

func Connect() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("pgx", connectionString)

	if err != nil {
		return nil, err
	}

	return db, err
}
