package repository

import (
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
)

var _ payments.Repository = (*PaymentRepository)(nil)

type PaymentRepository struct {
	psqlDB connection.DB
}

func NewPaymentRepository(psqlDB connection.DB) *PaymentRepository {
	return &PaymentRepository{psqlDB: psqlDB}
}
