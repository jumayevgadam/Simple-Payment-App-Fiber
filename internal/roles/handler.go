package roles

import "github.com/gofiber/fiber/v2"

// Handler interface for performing roles crud in this layer
type Handlers interface {
	AddRole() fiber.Handler
	GetRole() fiber.Handler
	GetRoles() fiber.Handler
	UpdateRole() fiber.Handler
	DeleteRole() fiber.Handler
}
