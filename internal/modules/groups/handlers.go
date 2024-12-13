package groups

import "github.com/gofiber/fiber/v2"

// Handler interface for groups.
type Handlers interface {
	AddGroup() fiber.Handler
	GetGroup() fiber.Handler
	ListGroups() fiber.Handler
	DeleteGroup() fiber.Handler
	UpdateGroup() fiber.Handler
	ListStudentsByGroupID() fiber.Handler
}
