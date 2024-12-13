package payment

import "github.com/gofiber/fiber/v2"

// Handlers interface for performing payment operations.
type Handlers interface {
	AddPayment() fiber.Handler
	GetPaymentByID() fiber.Handler
	StudentListPaymentsByStudentID() fiber.Handler
	AdminListPaymentsByStudentID() fiber.Handler
	UpdatePaymentByStudent() fiber.Handler
	ChangePaymentStatus() fiber.Handler
}
