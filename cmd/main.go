package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/tsu-toleg/internal/common/metrics"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/cache"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database/postgres"
	"github.com/jumayevgadam/tsu-toleg/internal/server"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
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

	// if our server stops to run then metrics will not affect from server, it will run forever.
	go func() {
		err = metrics.Listen(cfg.Server.MetricsPort)
		log.Printf("[metrics][Listen]: %v", err)
	}()

	// Redis Connection is
	var rdb connection.Cache
	rdb, err = connection.NewCache(ctx, cfg.Redis)
	if err != nil {
		log.Printf("connection.NewCache in main: %v", err.Error())
	}

	defer func() {
		err = rdb.Close()
		if err != nil {
			log.Printf("redis.Close in main: %v", err)
		}
	}()

	dataStore := postgres.NewDataStore(psqlDB)
	cacheStore := cache.NewClientRDRepository(rdb)

	source := server.NewServer(cfg, dataStore, cacheStore, appLogger)
	if err := source.Run(); err != nil {
		log.Printf("error in runnning application: %v", err.Error())
	}
}
