package main

import (
	"context"
	"log"

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
	err := application.BootStrap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
