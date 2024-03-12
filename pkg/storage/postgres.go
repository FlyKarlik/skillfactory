package storage

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresDb(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	//defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
