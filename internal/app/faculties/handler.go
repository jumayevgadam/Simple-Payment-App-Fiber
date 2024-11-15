package faculties

import "github.com/gofiber/fiber/v2"

// Handlers interface for faculties.
type Handlers interface {
	AddFaculty() fiber.Handler
	GetFaculty() fiber.Handler
	ListFaculties() fiber.Handler
	DeleteFaculty() fiber.Handler
	UpdateFaculty() fiber.Handler
}
