package payments

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	AddPayment() fiber.Handler
	StudentUpdatePayment() fiber.Handler
	GetPayment() fiber.Handler
	ListPaymentsByStudent() fiber.Handler
	StudentDeletePayment() fiber.Handler

	AdminListPaymentsByStudent() fiber.Handler
	AdminUpdatePaymentOfStudent() fiber.Handler
	AdminDeleteStudentPayment() fiber.Handler

	// ------- STATISTICS ------------------//.

	AdminGetStatisticsAboutYear() fiber.Handler
	AdminGetStatisticsAboutFaculty() fiber.Handler
}
