package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	userOps "github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
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

// CreateUser handler creates a new user and returns id.
// @Summary Register User.
// @Description create user func general func for creating users.
// @Tags Users
// @ID register-user
// @Accept multipart/form-data
// @Produce json
// @Param req formData userModel.SignUpReq true "register user payload"
// @Success 200 {int} int
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /auth/register [post].
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
// @Summary Login
// @Description login func for all users.
// @Tags Users
// @ID login
// @Accept multipart/form-data
// @Produce json
// @Param role path string true "role"
// @Param loginReq formData userModel.LoginReq true "login request payload"
// @Success 200 {object} string
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /auth/{role}/login [post].
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

// ChangeRoleOfUser handler.
func (h *UserHandler) ChangeRoleOfUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// role, ok := c.Locals("role_type").(string)
		// if !ok || role != "superadmin" {
		// 	return errlst.NewUnauthorizedError("only superadmin can see list of all users")
		// }

		userID, err := strconv.Atoi(c.Params("user_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		err = h.service.UserService().UpdateUser(c.Context(), userID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON("user's role successfully changed")
	}
}

// ListUsers handler.
func (h *UserHandler) ListUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != "superadmin" {
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
		return nil
	}
}

// UpdateUser handler.
func (h *UserHandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// GetUser handler.
func (h *UserHandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
