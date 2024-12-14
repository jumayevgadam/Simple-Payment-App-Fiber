package users

import "github.com/gofiber/fiber/v2"

// Handlers interface for users.
type Handlers interface {
	// Admin.
	AddAdmin() fiber.Handler
	GetAdmin() fiber.Handler
	ListAdmins() fiber.Handler
	DeleteAdmin() fiber.Handler

	AddStudent() fiber.Handler
	GetStudent() fiber.Handler
	ListStudents() fiber.Handler
	DeleteStudent() fiber.Handler
	// Student.
}
