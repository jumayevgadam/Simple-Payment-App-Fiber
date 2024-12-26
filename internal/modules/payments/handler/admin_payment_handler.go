package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (h *PaymentHandler) GetPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := middleware.GetStudentIDFromFiberContext(c)
		if err != nil {
			return errlst.Response(c, errlst.NewUnauthorizedError(err.Error()))
		}

		paymentIDStr := c.Query("payment-id")
		if paymentIDStr == "" {
			return errlst.Response(c, errlst.NewBadQueryParamsError("payment-id must implement in query param"))
		}

		paymentID, err := strconv.Atoi(paymentIDStr)
		if err != nil {
			return errlst.Response(c, errlst.NewBadRequestError(err.Error()))
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

		listPaymentsByStudent, studentRes, err := h.service.PaymentService().ListPaymentsByStudent(c.Context(), studentID, paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"studentPayments": listPaymentsByStudent,
			"student":         studentRes,
		})
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

func (h *PaymentHandler) AdminDeleteStudentPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := strconv.Atoi(c.Params("student_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		paymentID, err := strconv.Atoi(c.Params("payment_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.PaymentService().AdminDeleteStudentPayment(c.Context(), studentID, paymentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "admin successfully deleted payments of student",
		})
	}
}
