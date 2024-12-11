package main

import (
	"context"
	"log"

	_ "github.com/jumayevgadam/tsu-toleg/docs"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/tsu-toleg/internal/server"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// @title TSU-TOLEG API Documentation
// @version 2.0
// @description This is the API for the TSU-TOLEG system.
// @termsOfService http://swagger.io/terms
// @contact.name Gadam Jumayev
// @contact.url https://github.com/jumayevgadam
// @contact.email hypergadam@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /api/v1
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

	// INITIALIZE DATASTORE.
	dataStore := postgres.NewDataStore(psqlDB)

	// INITIALIZE SERVER AND HANDLERS.
	srv := server.NewServer(cfg, dataStore, appLogger)
	srv.MapHandlers(dataStore)

	// INITIALIZE SRV.RUN() WITH GOROUTINE.
	if err := srv.Fiber.Listen(":" + cfg.Server.HTTPPort); err != nil {
		appLogger.Errorf("error listening on http port: %v", err.Error())
	}

	// HANDLE GRACEFUL SHUTDOWN.
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// This blocks the main thread until an interrupt is received.
	// <-quit

	// gracefulCtx, cancel := context.WithTimeout(context.Background(), constants.CtxDefaultTimeOut)
	// defer cancel()

	if err := srv.Fiber.ShutdownWithContext(context.Background()); err != nil {
		appLogger.Errorf("error occured when shutting down application")
	}

	appLogger.Info("application successfully shutdown.")

	// CLOSE psqlDB.
	if err := psqlDB.Close(); err != nil {
		appLogger.Errorf("error closing psqlDB: %v", err.Error())
	}
}
