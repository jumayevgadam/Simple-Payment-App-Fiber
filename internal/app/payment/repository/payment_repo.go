package repository

import (
	"github.com/jumayevgadaym/tsu-toleg/internal/app/payment"
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
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
