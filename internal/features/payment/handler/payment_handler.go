package handler

import paymentOps "github.com/jumayevgadaym/tsu-toleg/internal/app/payment"

// Ensure PaymentHandler implements the paymentOps.Handler interface.
var (
	_ paymentOps.Handler = (*PaymentHandler)(nil)
)

// PaymentHandler for working with http requests.
type PaymentHandler struct {
	service paymentOps.Service
}

// NewPaymentHandler creates and returns a new instance of PaymentHandler.
func NewPaymentHandler(service paymentOps.Service) *PaymentHandler {
	return &PaymentHandler{service: service}
}
