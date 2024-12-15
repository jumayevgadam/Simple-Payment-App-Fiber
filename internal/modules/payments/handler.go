package payments

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	AddPayment() fiber.Handler
	UpdatePayment() fiber.Handler
	GetPayment() fiber.Handler
	ListPaymentsByStudent() fiber.Handler

	AdminListPaymentsByStudent() fiber.Handler
}
