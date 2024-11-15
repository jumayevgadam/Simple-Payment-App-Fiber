package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/users/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/users/service"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/internal/middleware/token"
)

// Routes function for users in this place
func Routes(f fiber.Router, dataStore database.DataStore) {
	tokenGenerator := token.NewTokenOps(config.JWTOps{})
	// Init Service
	Service := service.NewUserService(tokenGenerator, dataStore)
	// Init Handler
	Handler := handler.NewUserHandler(&config.Config{}, Service)

	// groups
	// adminGroup is
	adminGroup := f.Group("/admin")
	{
		adminGroup.Post("/sign-up", Handler.CreateUser())
	}

	// userGroup is
	userGroup := f.Group("/user")
	{
		userGroup.Post("/sign-up", Handler.CreateUser())
		userGroup.Post("/login", Handler.Login())
	}
}
