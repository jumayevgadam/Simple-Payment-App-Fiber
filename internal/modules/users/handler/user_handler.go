package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	userOps "github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

// Ensure UserHandler implements the userOps.Handler interface.
var (
	_ userOps.Handler = (*UserHandler)(nil)
)

// UserHandler manages http request methods and calls methods from service and config.
type UserHandler struct {
	mw      *middleware.MiddlewareManager
	service userOps.Service
}

// NewUserHandler creates and returns a new instance of UserHandler.
func NewUserHandler(mw *middleware.MiddlewareManager, service userOps.Service) *UserHandler {
	return &UserHandler{mw: mw, service: service}
}

// CreateUser handler creates a new user and returns id.
// @Summary Create User.
// @Description create user func general func for creating users.You can use for role superadmin, admin, student.
// @Tags Users
// @ID create-user
// @Accept multipart/form-data
// @Produce json
// @Param role path string true "role"
// @Param req formData userModel.SignUpReq true "create user payload"
// @Success 200 {int} int
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /auth/{role}/sign-up [post]
func (h *UserHandler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Params("role")

		var req userModel.SignUpReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			return errlst.Response(c, err)
		}

		userID, err := h.service.CreateUser(c.Context(), req, role)
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
// @Success 200 {object} userModel.UserWithTokens
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /auth/{role}/login [post]
func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Params("role")

		var loginReq userModel.LoginReq
		if err := reqvalidator.ReadRequest(c, &loginReq); err != nil {

			return errlst.Response(c, err)
		}

		userWithToken, err := h.service.Login(c.Context(), loginReq, role)
		if err != nil {
			return errlst.Response(c, err)
		}

		utils.SetAuthCookies(c, userWithToken.AccessToken, userWithToken.RefreshToken)

		return c.Status(fiber.StatusOK).JSON(userWithToken)
	}
}
