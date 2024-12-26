package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"

	paymentModel "github.com/jumayevgadam/tsu-toleg/internal/models/payment"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

func (h *PaymentHandler) AddPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		var request paymentModel.Request

		err = reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		checkPhoto, err := utils.ReadImage(c, "check-photo")
		if err != nil {
			return errlst.Response(c, err)
		}

		paymentID, err := h.service.PaymentService().AddPayment(c, checkPhoto, studentID, request)
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

		listPaymentsByStudent, _, err := h.service.PaymentService().ListPaymentsByStudent(c.Context(), studentID, paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listPaymentsByStudent)
	}
}

func (h *PaymentHandler) StudentDeletePayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.Response(c, errlst.NewUnauthorizedError(err.Error()))
		}

		paymentID, err := strconv.Atoi(c.Params("payment_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.PaymentService().StudentDeletePayment(c.Context(), studentID, paymentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "payment successfully deleted",
		})
	}
}
