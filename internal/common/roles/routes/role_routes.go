package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/cache"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/roles/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/roles/service"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
)

// Routes is
func Routes(f fiber.Router, dataStore database.DataStore, cacheStore cache.Store) {
	// init service
	Service := service.NewRoleService(dataStore, cacheStore)
	// init handler
	Handler := handler.NewRoleHandler(Service)

	// roleGroup is
	roleGroup := f.Group("/role")
	{
		roleGroup.Post("/create", Handler.AddRole())
		roleGroup.Get("/get-all", Handler.GetRoles())
		roleGroup.Get("/:id", Handler.GetRole())
		roleGroup.Delete("/:id", Handler.DeleteRole())
		roleGroup.Put("/:id", Handler.UpdateRole())
	}
}
