package main

import (
	"context"
	"log"

	"github.com/jumayevgadaym/tsu-toleg/internal/cache"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/database/postgres"
	"github.com/jumayevgadaym/tsu-toleg/internal/metrics"
	"github.com/jumayevgadaym/tsu-toleg/internal/server"
	"github.com/jumayevgadaym/tsu-toleg/pkg/constants"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeout)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error in main.LoadConfig: %v", err.Error())
	}

	// PostgreSQL connection
	psqlDB, err := connection.GetDBConnection(ctx, cfg.Postgres)
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

	// if our server stops to run then metrics will not affect from server, it will run forever
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
	log.Println("redisDB connected!!")

	defer func() {
		err = rdb.Close()
		if err != nil {
			log.Printf("redis.Close in main: %v", err)
		}
	}()

	dataStore := postgres.NewDataStore(psqlDB)
	cacheStore := cache.NewClientRDRepository(rdb)

	source := server.NewServer(cfg, dataStore, cacheStore)
	if err := source.Run(); err != nil {
		log.Printf("error in runnning application: %v", err.Error())
	}
}
