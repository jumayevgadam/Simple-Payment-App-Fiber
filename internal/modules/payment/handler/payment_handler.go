package handler

import (
	"github.com/gofiber/fiber/v2"
	paymentOps "github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
)

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

// AddPayment for students.
func (h *PaymentHandler) AddPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// GetPayment About student.
func (h *PaymentHandler) GetPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// UpdatePayment details.
func (h *PaymentHandler) UpdatePayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// ListPayments using pagination.
func (h *PaymentHandler) ListPayments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// GetPaymentByStudent ops.
func (h *PaymentHandler) GetPaymentByStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
