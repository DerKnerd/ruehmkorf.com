package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return db, err
}
