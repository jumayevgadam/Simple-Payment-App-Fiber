package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/helpers"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payments"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

var _ payments.Handlers = (*PaymentHandler)(nil)

type PaymentHandler struct {
	service services.DataService
}

func NewPaymentHandler(service services.DataService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) AddPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		var request paymentModel.Request

		err = reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		if err := helpers.CheckPayment(request); err != nil {
			return errlst.NewBadRequestError(err)
		}

		checkPhoto, err := utils.ReadImage(c, "check-photo")
		if err != nil {
			return errlst.Response(c, err)
		}

		paymentID, err := h.service.PaymentService().AddPayment(c, checkPhoto, studentID, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "payment successfully added",
			"paymentID": paymentID,
		})
	}
}

func (h *PaymentHandler) StudentUpdatePayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.NewUnauthorizedError(err)
		}

		paymentID, err := strconv.Atoi(c.Params("payment_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var updateRequest paymentModel.UpdatePaymentRequest

		err = reqvalidator.ReadRequest(c, &updateRequest)
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		res, err := h.service.PaymentService().StudentUpdatePayment(c, studentID, paymentID, updateRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func (h *PaymentHandler) GetPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.NewUnauthorizedError(err)
		}

		paymentIDStr := c.Query("payment-id")
		if paymentIDStr == "" {
			return errlst.NewBadRequestError("[paymentHandler][GetPayment]: payment-id Query param must need")
		}

		paymentID, err := strconv.Atoi(paymentIDStr)
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		paymentRes, err := h.service.PaymentService().GetPayment(c.Context(), studentID, paymentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"message":  "successfully got payment",
				"response": paymentRes,
			},
		)
	}
}

func (h *PaymentHandler) ListPaymentsByStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.NewUnauthorizedError(err)
		}

		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err)
		}

		listPaymentsByStudent, err := h.service.PaymentService().ListPaymentsByStudent(c.Context(), studentID, paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listPaymentsByStudent)
	}
}

func (h *PaymentHandler) AdminListPaymentsByStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentIDStr := c.Query("student-id")
		if studentIDStr == "" {
			return errlst.NewBadQueryParamsError("student-id param cannot be empty")
		}

		studentID, err := strconv.Atoi(studentIDStr)
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err)
		}

		listPaymentsByStudent, err := h.service.PaymentService().ListPaymentsByStudent(c.Context(), studentID, paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listPaymentsByStudent)
	}
}

func (h *PaymentHandler) AdminUpdatePaymentOfStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := strconv.Atoi(c.Params("student_id"))
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		paymentID, err := strconv.Atoi(c.Params("payment_id"))
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		paymentStatus := c.FormValue("payment-status")

		response, err := h.service.PaymentService().AdminUpdatePaymentOfStudent(
			c.Context(), studentID, paymentID, paymentStatus)

		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "success in update",
			"response": response,
		})
	}
}
