package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	"github.com/jumayevgadaym/tsu-toleg/internal/faculties/handler"
	"github.com/jumayevgadaym/tsu-toleg/internal/faculties/service"
)

// Routes func for faculty routes
func Routes(f fiber.Router, dataStore database.DataStore) {
	// Init Service
	Service := service.NewFacultyService(dataStore)
	// Init Handler
	Handler := handler.NewFacultyHandler(Service)

	// facultyGroup is
	facultyGroup := f.Group("/faculty")
	{
		facultyGroup.Post("/create", Handler.AddFaculty())
		facultyGroup.Get("/get-all", Handler.ListFaculties())
		facultyGroup.Get("/:id", Handler.GetFaculty())
		facultyGroup.Delete("/:id", Handler.DeleteFaculty())
		facultyGroup.Put("/:id", Handler.UpdateFaculty())
	}
}
