package routes

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users/service"
)

// Routes function for users in this place.
func Routes(f fiber.Router, mw *mwMngr.MiddlewareManager, dataStore database.DataStore) {
	// Init Service.
	Service := service.NewUserService(mw, dataStore)
	// Init Handler.
	Handler := handler.NewUserHandler(mw, Service)

	authGroup := f.Group("/auth")
	{
		authGroup.Post("/register", Handler.Register())
		authGroup.Post("/login", Handler.Login())
		// authGroup.Post("/:role/logout", Handler.Logout())
	}
}
