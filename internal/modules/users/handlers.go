package users

import "github.com/gofiber/fiber/v2"

// Handlers interface for users.
type Handlers interface {
	// General.
	Login() fiber.Handler
	// Admin.
	AddAdmin() fiber.Handler
	GetAdmin() fiber.Handler
	ListAdmins() fiber.Handler
	DeleteAdmin() fiber.Handler
	UpdateAdmin() fiber.Handler
	AdminFindStudent() fiber.Handler

	AddStudent() fiber.Handler
	GetStudent() fiber.Handler
	ListStudents() fiber.Handler
	DeleteStudent() fiber.Handler
	UpdateStudent() fiber.Handler
	ListPaymentsByStudentID() fiber.Handler
}
