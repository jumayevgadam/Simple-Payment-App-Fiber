package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/tsu-toleg/internal/server"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeout)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error in main.LoadConfig: %v", err.Error())
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("Mode: %s", cfg.Server.Mode)

	// PostgreSQL connection
	psqlDB, err := connection.GetDBConnection(ctx, cfg.Postgres)
	if err != nil {
		log.Printf("error in getting DB connection: %v", err.Error())
	}

	defer func() {
		if err := psqlDB.Close(); err != nil {
			log.Printf("error in closing DB: %v", err.Error())
		}
	}()

	dataStore := postgres.NewDataStore(psqlDB)

	source := server.NewServer(cfg, dataStore, appLogger)
	if err := source.Run(); err != nil {
		log.Printf("error in runnning application: %v", err.Error())
	}
}
