package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	paymentOps "github.com/jumayevgadam/tsu-toleg/internal/modules/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

// Ensure PaymentHandler implements the paymentOps.Handler interface.
var (
	_ paymentOps.Handlers = (*PaymentHandler)(nil)
)

// PaymentHandler for working with http requests.
type PaymentHandler struct {
	service services.DataService
}

// NewPaymentHandler creates and returns a new instance of PaymentHandler.
func NewPaymentHandler(service services.DataService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

// AddPayment for students.
func (h *PaymentHandler) AddPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals(constants.RoleType).(string)
		if !ok || role != constants.Student {
			return errlst.NewUnauthorizedError("only student role can perform payment")
		}

		studentID, ok := c.Locals(constants.UserID).(int)
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

		paymentID, err := h.service.PaymentService().AddPayment(c, studentID, checkPhoto, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(paymentID)
	}
}

// GetPayment About student.
func (h *PaymentHandler) GetPaymentByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paymentID, err := strconv.Atoi(c.Params("payment_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		paymentDTO, err := h.service.PaymentService().GetPaymentByID(c.Context(), paymentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(paymentDTO)
	}
}

// GetPaymentByStudentID handler.
func (h *PaymentHandler) StudentListPaymentsByStudentID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != "student" {
			return errlst.NewUnauthorizedError("only student can see his(her) payments")
		}

		studentID, ok := c.Locals("user_id").(int)
		if !ok {
			return errlst.NewUnauthorizedError("error in type assertion:[StudentListPaymentsByStudentID]")
		}

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err.Error())
		}

		studentPayments, err := h.service.PaymentService().StudentListPaymentsByStudentID(
			c.Context(), studentID, paginationReq)

		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(studentPayments)
	}
}

func (h *PaymentHandler) UpdatePaymentByStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (h *PaymentHandler) ChangePaymentStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (h *PaymentHandler) AdminListPaymentsByStudentID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
