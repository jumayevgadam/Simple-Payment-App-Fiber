package payments

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	AddPayment() fiber.Handler
	StudentUpdatePayment() fiber.Handler
	GetPayment() fiber.Handler
	ListPaymentsByStudent() fiber.Handler

	AdminListPaymentsByStudent() fiber.Handler
}
