package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/groups/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/app/groups/service"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
)

// Routes for groups.
func Routes(f fiber.Router, dataStore database.DataStore) {
	// Init Service.
	Service := service.NewGroupService(dataStore)
	// Init Handler.
	Handler := handler.NewGroupHandler(Service)
	// put route.
	groupPath := f.Group("/group")
	{
		groupPath.Post("/add", Handler.AddGroup())
		groupPath.Get("/get-all", Handler.ListGroups())
		groupPath.Get("/:id", Handler.GetGroup())
		groupPath.Delete("/:id", Handler.DeleteGroup())
		groupPath.Put("/:id", Handler.UpdateGroup())
	}
}
