package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

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

func (h *UserHandler) UpdateAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminID, err := strconv.Atoi(c.Params("admin_id"))
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		var updateRequest userModel.AdminUpdateRequest

		err = reqvalidator.ReadRequest(c, &updateRequest)
		if err != nil {
			res, err := updateRequest.Validate()
			if err == nil {
				return c.Status(fiber.StatusOK).JSON(res)
			}

			return errlst.Response(c, err)
		}

		updateResponse, err := h.service.UserService().UpdateAdmin(c.Context(), adminID, updateRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": updateResponse,
		})
	}
}

func (h *UserHandler) AdminFindStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filterStudent := userModel.GetQueryParamsForFilterStudents(c)

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err)
		}

		studentListWithFilter, err := h.service.UserService().AdminFindStudent(c.Context(), filterStudent, paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(studentListWithFilter)
	}
}
