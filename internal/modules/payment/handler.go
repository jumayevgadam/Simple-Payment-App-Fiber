package payment

import "github.com/gofiber/fiber/v2"

// Handler interface for performing payment operations.
type Handlers interface {
	AddPayment() fiber.Handler
	UpdatePayment() fiber.Handler
	ListPayments() fiber.Handler
	GetPaymentByStudent() fiber.Handler
}
