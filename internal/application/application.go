package application

import (
	"context"
	"log"
	"runtime"

	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/tsu-toleg/internal/server"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// BootStrap application.
func BootStrap(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error in main.LoadConfig: %v", err.Error())
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("Mode: %s", cfg.Server.Mode)

	numCPUs := runtime.NumCPU()
	appLogger.Infof("Number of CPU's: %d", numCPUs)

	runtime.GOMAXPROCS(numCPUs)

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

	appLogger.Info("Server Started\n")

	err = source.Run()
	if err != nil {
		return errlst.ParseErrors(err)
	}

	appLogger.Info("Server stopped gracefully...\n")
	return nil
}
