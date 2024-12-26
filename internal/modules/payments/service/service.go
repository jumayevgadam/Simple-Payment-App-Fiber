package service

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
)

var _ payments.Service = (*PaymentService)(nil)

type PaymentService struct {
	repo database.DataStore
}

func NewPaymentService(repo database.DataStore) *PaymentService {
	return &PaymentService{repo: repo}
}
