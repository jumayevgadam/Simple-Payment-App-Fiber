package application

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	serverErrors := make(chan error, 1)
	go func() {
		if err := source.Run(); err != nil {
			serverErrors <- err
		}
	}()
	appLogger.Info("Server Started\n")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		appLogger.Infof("Caught Signal: %v, graceful shutdown....\n", sig)
	case err := <-serverErrors:
		appLogger.Errorf("Server error: %v\n", err.Error())
		return errlst.ParseErrors(err)
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.Server.CtxDefaultTimeOut)
	defer cancel()

	if err := source.Stop(ctx); err != nil {
		appLogger.Errorf("Can not stop server")
		return errlst.ParseErrors(err)
	}

	appLogger.Info("Server stopped gracefully...\n")
	return nil
}
