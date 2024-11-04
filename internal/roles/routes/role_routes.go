package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles/service"
)

// Routes is
func Routes(f fiber.Router, dataStore database.DataStore) {
	// init service
	Service := service.NewRoleService(dataStore)
	// init handler
	Handler := handler.NewRoleHandler(Service)

	// roleGroup is
	roleGroup := f.Group("/role")
	{
		roleGroup.Post("/create", Handler.AddRole())
		roleGroup.Get("/:id", Handler.GetRole())
	}
}
