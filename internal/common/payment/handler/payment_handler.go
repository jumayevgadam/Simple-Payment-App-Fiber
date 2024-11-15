package handler

import paymentOps "github.com/jumayevgadaym/tsu-toleg/internal/payment"

var (
	_ paymentOps.Handler = (*PaymentHandler)(nil)
)

// PaymentHandler is
type PaymentHandler struct {
	service paymentOps.Service
}

// NewPaymentHandler is
func NewPaymentHandler(service paymentOps.Service) *PaymentHandler {
	return &PaymentHandler{service: service}
}
