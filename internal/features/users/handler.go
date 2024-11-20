package users

import "github.com/gofiber/fiber/v2"

// Handler interface for performing users operations.
type Handler interface {
	CreateUser() fiber.Handler
	Login(role string) fiber.Handler
	// ListUsers() fiber.Handler
	// DeleteUser() fiber.Handler
	// UpdateUser() fiber.Handler
	// GetUser() fiber.Handler
}
