package payment

import (
	"context"

	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for performing payment actions in repo layer.
type Repository interface {
	AddPayment(ctx context.Context, data *paymentModel.Response) (int, error)
	GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDAO, error)
	CountPaymentsByStudentID(ctx context.Context, studentID int) (int, error)

	StudentListPaymentsByStudentID(ctx context.Context, studentID int, paginationData abstract.PaginationData) (
		[]*paymentModel.AllPaymentDAO, error,
	)
}
