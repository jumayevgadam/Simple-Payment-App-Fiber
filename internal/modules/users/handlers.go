package users

import "github.com/gofiber/fiber/v2"

// Handler interface for performing users operations.
type Handlers interface {
	Register() fiber.Handler
	Login() fiber.Handler
	ListUsers() fiber.Handler
	ChangeRoleOfUser() fiber.Handler
	DeleteUser() fiber.Handler
	UpdateUser() fiber.Handler
	GetUser() fiber.Handler
}
