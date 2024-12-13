package users

import "github.com/gofiber/fiber/v2"

// Handlers interface for performing users operations.
type Handlers interface {
	Register() fiber.Handler
	Login() fiber.Handler
	ListUsers() fiber.Handler
	DeleteUser() fiber.Handler
	UpdateUser() fiber.Handler
	GetUserByID() fiber.Handler

	ListStudents() fiber.Handler
	// ListStudentsByFaculty() fiber.Handler
}
