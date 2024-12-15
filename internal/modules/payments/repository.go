package payments

import (
	"context"

	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

type Repository interface {
	AddPayment(ctx context.Context, paymentData *paymentModel.Response) (int, error)
	GetStudentInfoForPayment(ctx context.Context, studentID int) (*paymentModel.StudentInfoForPayment, error)
	GetPaymentByID(ctx context.Context, paymentID int) (*paymentModel.AllPaymentDAO, error)
	CountPaymentByStudent(ctx context.Context, studentID int) (int, error)
	ListPaymentsByStudent(ctx context.Context, studentID int, paginationData abstract.PaginationData) (
		[]*paymentModel.AllPaymentDAO, error,
	)
}
