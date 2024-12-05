package payment

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
)

// Service interface for performing payment actions in service layer.
type Service interface {
	AddPayment(ctx *fiber.Ctx, studentID int, checkPhoto *multipart.FileHeader, req *paymentModel.Request) (int, error)
	// GetPaymentDetails, use view
	// UpdatePayment, check payment
	// ListPaymentOfStudent
	// CheckPayment Status

	// ````````- STATISTICS -````````
	// GetStatisticsOfStudents
}
