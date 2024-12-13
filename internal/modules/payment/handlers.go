package payment

import "github.com/gofiber/fiber/v2"

// Handlers interface for performing payment operations.
type Handlers interface {
	AddPayment() fiber.Handler
	GetPaymentByID() fiber.Handler
	GetPaymentByStudentID() fiber.Handler
}
