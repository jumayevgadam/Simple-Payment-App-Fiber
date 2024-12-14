package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

var _ users.Handlers = (*UserHandler)(nil)

type UserHandler struct {
	service services.DataService
}

func NewUserHandler(service services.DataService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) AddStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request userModel.Request

		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
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

func (h *UserHandler) AddAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request userModel.AdminRequest

		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		adminID, err := h.service.UserService().AddAdmin(c.Context(), request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "admin successfully created",
			"adminID": adminID,
		})
	}
}

func (h *UserHandler) ListAdmins() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err.Error())
		}

		adminList, err := h.service.UserService().ListAdmins(c.Context(), paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(adminList)
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

func (h *UserHandler) GetAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminID, err := strconv.Atoi(c.Params("admin_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		admin, err := h.service.UserService().GetAdmin(c.Context(), adminID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(admin)
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

func (h *UserHandler) DeleteAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminID, err := strconv.Atoi(c.Params("admin_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		err = h.service.UserService().DeleteAdmin(c.Context(), adminID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
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
