package main

import (
	"learn/go/cmd/api"
	"learn/go/config"
	"learn/go/db"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	pgxConfig, err := pgx.ParseConfig(config.Envs.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error while parsing database url")
	}
	db, err := db.NewPGStorage(*pgxConfig)
	if err != nil {
		log.Fatalf("Error connecting to DB %v", err)
	}
	server := api.NewAPIServer(":8080", db)
	err = server.Run()
	if err != nil {
		log.Fatal("Error strating the server")
	}

}
