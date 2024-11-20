package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/middleware"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/middleware/token"
	"github.com/jumayevgadaym/tsu-toleg/internal/features/groups/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/features/groups/service"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
)

// Routes for groups.
func Routes(f fiber.Router, tp *token.TokenOps, dataStore database.DataStore) {
	// Init Service.
	Service := service.NewGroupService(dataStore)
	// Init Handler.
	Handler := handler.NewGroupHandler(Service)
	// put route.
	groupPath := f.Group("/group")
	// groupPath.Use(middleware.RoleBasedMiddleware(cfg.JWT, 1, 2, 3))
	{
		groupPath.Post("/add", Handler.AddGroup())
		groupPath.Get("/get-all", Handler.ListGroups())
		groupPath.Get("/:id", middleware.RoleBasedMiddleware(tp, 3), Handler.GetGroup())
		groupPath.Delete("/:id", Handler.DeleteGroup())
		groupPath.Put("/:id", Handler.UpdateGroup())
	}
}
