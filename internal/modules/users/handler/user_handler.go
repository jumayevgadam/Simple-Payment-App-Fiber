package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	userOps "github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// Ensure UserHandler implements the userOps.Handler interface.
var (
	_ userOps.Handlers = (*UserHandler)(nil)
)

// UserHandler manages http request methods and calls methods from service and config.
type UserHandler struct {
	service services.DataService
}

// NewUserHandler creates and returns a new instance of UserHandler.
func NewUserHandler(service services.DataService) *UserHandler {
	return &UserHandler{service: service}
}

// Register handler creates a new user and returns id.
func (h *UserHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request userModel.SignUpReq

		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		userID, err := h.service.UserService().Register(c.Context(), request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(userID)
	}
}

// Login handler method for login.
func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginReq userModel.LoginReq

		err := reqvalidator.ReadRequest(c, &loginReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		token, err := h.service.UserService().Login(c.Context(), loginReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(token)
	}
}

// ListUsers handler.
func (h *UserHandler) ListUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != constants.SuperAdmin {
			return errlst.NewUnauthorizedError("only superadmin can see list of all users")
		}

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		listOfUsers, err := h.service.UserService().ListAllUsers(c.Context(), paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listOfUsers)
	}
}

// DeleteUser handler.
func (h *UserHandler) DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != constants.SuperAdmin {
			return errlst.NewUnauthorizedError("only superadmin can delete user.")
		}

		userID, err := strconv.Atoi(c.Params("user_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		err = h.service.UserService().DeleteUser(c.Context(), userID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON("user deleted successfully")
	}
}

// UpdateUser handler.
func (h *UserHandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != constants.SuperAdmin {
			return errlst.NewUnauthorizedError("only superadmin can update user details.")
		}

		userID, err := strconv.Atoi(c.Params("user_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		var updateReq userModel.UpdateUserDetails

		err = reqvalidator.ReadRequest(c, &updateReq)
		if err != nil {
			noUpdateRes, err := updateReq.Validate()
			if err == nil {
				return c.Status(fiber.StatusOK).JSON(noUpdateRes)
			}

			return errlst.NewBadRequestError(err.Error())
		}

		err = h.service.UserService().UpdateUser(c.Context(), userID, &updateReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON("successfully updated user details.")
	}
}

// GetUserByID handler.
func (h *UserHandler) GetUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("user_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		userRes, err := h.service.UserService().GetUserByID(c.Context(), userID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(userRes)
	}
}

// ListStudents handler.
func (h *UserHandler) ListStudents() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		listOfStudents, err := h.service.UserService().ListStudents(c.Context(), paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listOfStudents)
	}
}
