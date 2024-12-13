package payment

import (
	"context"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Service interface for performing payment actions in service layer.
type Service interface {
	AddPayment(ctx *fiber.Ctx, studentID int, checkPhoto *multipart.FileHeader, req *paymentModel.Request) (int, error)
	GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDTO, error)

	StudentListPaymentsByStudentID(ctx context.Context, studentID int, paginationQuery abstract.PaginationQuery) (
		abstract.PaginatedResponse[*paymentModel.AllPaymentDTO], error,
	)
}
