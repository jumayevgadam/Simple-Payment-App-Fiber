package times

import "github.com/gofiber/fiber/v2"

// Handlers interface for times.
type Handlers interface {
	AddTime() fiber.Handler
	GetTime() fiber.Handler
	// ListTimes() fiber.Handler
	// UpdateTime() fiber.Handler
	// DeleteTime() fiber.Handler
}
