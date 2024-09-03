package db

import (
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

func NewPGStorage(cfg pgx.ConnConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.ConnString())
	if err != nil {
		log.Fatal("Error connecting to db\n", err)
	}
	return db, nil
}
