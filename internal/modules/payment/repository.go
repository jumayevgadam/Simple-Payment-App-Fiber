package payment

import (
	"context"

	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
)

// Repository interface for performing payment actions in repo layer.
type Repository interface {
	AddPayment(ctx context.Context, data *paymentModel.Response) (int, error)
}
