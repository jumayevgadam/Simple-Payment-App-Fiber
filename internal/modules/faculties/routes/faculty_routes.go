package routes

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties/service"
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
		facultyGroup.Post("/create", mwMngr.RoleBasedMiddleware(mw, "add:faculty", dataStore), Handler.AddFaculty())
		facultyGroup.Get("/get-all", mwMngr.RoleBasedMiddleware(mw, "list:faculties", dataStore), Handler.ListFaculties())
		facultyGroup.Get("/:id", mwMngr.RoleBasedMiddleware(mw, "get:faculty", dataStore), Handler.GetFaculty())
		facultyGroup.Delete("/:id", mwMngr.RoleBasedMiddleware(mw, "delete:faculty", dataStore), Handler.DeleteFaculty())
		facultyGroup.Put("/:id", mwMngr.RoleBasedMiddleware(mw, "update:faculty", dataStore), Handler.UpdateFaculty())
	}
}
