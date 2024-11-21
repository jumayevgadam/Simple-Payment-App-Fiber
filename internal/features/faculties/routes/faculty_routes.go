package routes

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadaym/tsu-toleg/internal/common/middleware"
	"github.com/jumayevgadaym/tsu-toleg/internal/features/faculties/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/features/faculties/service"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
)

// Routes func for faculty routes.
func Routes(f fiber.Router, mw *mwMngr.MiddlewareManager, dataStore database.DataStore) {
	// Init Service.
	Service := service.NewFacultyService(dataStore)
	// Init Handler.
	Handler := handler.NewFacultyHandler(Service)

	// facultyGroup is.
	facultyGroup := f.Group("/faculty")
	{
		facultyGroup.Post("/create", Handler.AddFaculty())
		facultyGroup.Get("/get-all", mwMngr.RoleBasedMiddleware(mw, 1, 2, 3), Handler.ListFaculties())
		facultyGroup.Get("/:id", Handler.GetFaculty())
		facultyGroup.Delete("/:id", Handler.DeleteFaculty())
		facultyGroup.Put("/:id", Handler.UpdateFaculty())
	}
}
