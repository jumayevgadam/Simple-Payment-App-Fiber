package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	paymentOps "github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
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
		if c.Locals("role") == nil && c.Locals("userRoleID") == nil && c.Locals("userID") == nil && c.Locals("username") == nil {
			return errlst.Response(c, errors.New("unauthorized access to this source"))
		}

		role, ok := c.Locals("role").(string)
		if !ok || role != "student" {
			return errlst.NewUnauthorizedError("only student role can perform payment")
		}

		studentID, ok := c.Locals("userID").(int)
		if !ok {
			return errlst.NewUnauthorizedError("cannot find student id in context")
		}

		var request paymentModel.Request
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errlst.Response(c, err)
		}

		checkPhoto, err := utils.ReadImage(c, "check-photo")
		if err != nil {
			return errlst.Response(c, err)
		}

		paymentID, err := h.service.AddPayment(c, studentID, checkPhoto, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(200).JSON(paymentID)
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
