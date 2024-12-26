package handler

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
)

var _ payments.Handlers = (*PaymentHandler)(nil)

type PaymentHandler struct {
	service services.DataService
}

func NewPaymentHandler(service services.DataService) *PaymentHandler {
	return &PaymentHandler{service: service}
}
