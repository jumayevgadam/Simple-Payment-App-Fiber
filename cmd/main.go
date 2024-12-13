package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/tsu-toleg/internal/server"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

func main() {
	// INITIALIZE CONFIG.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error in main.LoadConfig: %v", err.Error())
	}

	// INITIALIZE LOGGER.
	appLogger := logger.NewAPILogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("Mode: %s", cfg.Server.Mode)

	// INITIALIZE PSQLDB CONNECTION.
	psqlDB, err := connection.GetDBConnection(context.Background(), cfg.Postgres)
	if err != nil {
		appLogger.Errorf("[main][connection][GetDBConnection]: error: %v", err.Error())
	}

	defer func() {
		if psqlDB != nil {
			err = psqlDB.Close()
			if err != nil {
				appLogger.Errorf("error closing database: %v", err.Error())
			}
		}
	}()

	// INITIALIZE DATASTORE.
	dataStore := postgres.NewDataStore(psqlDB)

	// INITIALIZE SERVER.
	srv := server.NewServer(cfg, dataStore, appLogger)
	if err := srv.Run(); err != nil {
		appLogger.Errorf("error occured when running application")
	}
}
