package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/cache"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users/service"
)

// Routes function for users in this place.
func Routes(f fiber.Router, mw *mwMngr.MiddlewareManager, dataStore database.DataStore, redisStore cache.Store) {
	// Init Service.
	Service := service.NewUserService(mw, dataStore, redisStore)
	// Init Handler.
	Handler := handler.NewUserHandler(&config.Config{}, Service)

	authGroup := f.Group("/auth")
	{
		authGroup.Post("/:role/sign-up", Handler.CreateUser())
		authGroup.Post("/:role/login", Handler.Login())
	}
}
