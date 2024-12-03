package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure PaymentRepository implements the payment.Repository interface.
var (
	_ payment.Repository = (*PaymentRepository)(nil)
)

// PaymentRepository handles database operations related with payments.
type PaymentRepository struct {
	psqlDB connection.DB
}

// NewPaymentRepository creates and returns a new instance of PaymentRepository.
func NewPaymentRepository(psqlDB connection.DB) *PaymentRepository {
	return &PaymentRepository{psqlDB: psqlDB}
}

// AddPayment repo method insert payment details into payments table.
func (r *PaymentRepository) AddPayment(ctx context.Context, data *paymentModel.Response) (int, error) {
	var paymentID int

	err := r.psqlDB.QueryRow(
		ctx,
		addPaymentQuery,
		data.StudentID,
		data.PaymentType,
		data.PaymentStatus,
		data.CurrentPaidSum,
		data.CheckPhoto,
	).Scan(&paymentID)
	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return paymentID, nil
}
