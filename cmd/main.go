package main

import (
	"context"
	"log"

	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/database/postgres"
	"github.com/jumayevgadaym/tsu-toleg/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error in main.LoadConfig: %v", err.Error())
	}

	psqlDB, err := connection.GetDBConnection(context.Background(), cfg.Postgres)
	if err != nil {
		log.Printf("error in getting DB connection: %v", err.Error())
	} else {
		log.Println("POSTGRESQL:=>DB CONNECTED!")
	}
	defer func() {
		if err := psqlDB.Close(); err != nil {
			log.Printf("error in closing DB: %v", err.Error())
		}
	}()

	dataStore := postgres.NewDataStore(psqlDB)

	source := server.NewServer(cfg, dataStore)
	if err := source.Run(); err != nil {
		log.Printf("error in runnning application: %v", err.Error())
	}
}
