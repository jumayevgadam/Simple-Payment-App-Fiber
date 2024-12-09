package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jumayevgadam/tsu-toleg/docs"
	"github.com/jumayevgadam/tsu-toleg/internal/application"
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
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Start the application.
	server, err := application.BootStrap(ctx)
	if err != nil {
		log.Println("error during application bootstrap")
	}

	// Wait for termination signal.
	<-ctx.Done()

	log.Println("Shutdown signal received")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(gracefulCtx); err != nil {
		server.Logger.Errorf("error shutting down application: %v", err.Error())
	}

	server.Logger.Info("Application gracefully stopped")
}
