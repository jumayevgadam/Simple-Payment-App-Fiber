package repository

import (
	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/payment"
)

var (
	_ payment.Repository = (*PaymentRepository)(nil)
)

// PaymentRepository is
type PaymentRepository struct {
	psqlDB connection.DB
}

// NewPaymentRepository is
func NewPaymentRepository(psqlDB connection.DB) *PaymentRepository {
	return &PaymentRepository{psqlDB: psqlDB}
}
