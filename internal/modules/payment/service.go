package payment

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
)

// Service interface for performing payment actions in service layer.
type Service interface {
	AddPayment(ctx *fiber.Ctx, studentID int, checkPhoto *multipart.FileHeader, req *paymentModel.Request) (int, error)
	GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDTO, error)
}
