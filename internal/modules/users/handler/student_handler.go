package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

func (h *UserHandler) AddStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request userModel.Request

		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		userID, err := h.service.UserService().AddStudent(c.Context(), request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "student successfully created",
			"studentID": userID,
		})
	}
}

func (h *UserHandler) ListStudents() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err.Error())
		}

		studentList, err := h.service.UserService().ListStudents(c.Context(), paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(studentList)
	}
}

func (h *UserHandler) GetStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := strconv.Atoi(c.Params("student_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		student, err := h.service.UserService().GetStudent(c.Context(), studentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(student)
	}
}

func (h *UserHandler) DeleteStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := strconv.Atoi(c.Params("student_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		err = h.service.UserService().DeleteStudent(c.Context(), studentID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func (h *UserHandler) UpdateStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID, err := strconv.Atoi(c.Params("student_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var updateRequest userModel.StudentUpdateRequest

		err = reqvalidator.ReadRequest(c, &updateRequest)
		if err != nil {
			res, err := updateRequest.Validate()
			if err == nil {
				return c.Status(fiber.StatusOK).JSON(res)
			}

			return errlst.Response(c, err)
		}
								
		updateResponse, err := h.service.UserService().UpdateStudent(c.Context(), studentID, updateRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": updateResponse,
		})
	}
}
