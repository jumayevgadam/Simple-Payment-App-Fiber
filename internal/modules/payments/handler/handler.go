package handler

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
)

// Ensure payment.Handlers implements PaymentHandler.
var _ payments.Handlers = (*PaymentHandler)(nil)

type PaymentHandler struct {
	service services.DataService
}

// NewPaymentHandler creates and returns a new instance of PaymentHandler.
func NewPaymentHandler(service services.DataService) *PaymentHandler {
	return &PaymentHandler{service: service}
}
