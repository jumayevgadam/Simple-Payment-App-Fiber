package routes

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/groups/service"
)

// Routes for groups.
func Routes(f fiber.Router, mw *mwMngr.MiddlewareManager, dataStore database.DataStore) {
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
