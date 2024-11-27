package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	userOps "github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Ensure UserHandler implements the userOps.Handler interface.
var (
	_ userOps.Handler = (*UserHandler)(nil)
)

// UserHandler manages http request methods and calls methods from service and config.
type UserHandler struct {
	cfg     *config.Config
	service userOps.Service
}

// NewUserHandler creates and returns a new instance of UserHandler.
func NewUserHandler(cfg *config.Config, service userOps.Service) *UserHandler {
	return &UserHandler{cfg: cfg, service: service}
}

// CreateUser handler creates a new user and returns id.
// @Summary Create User.
// @Description create user func general func for creating users.
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
		ctx, span := otel.Tracer("[UserHandler]").Start(c.Context(), "[CreateUser]")
		defer span.End()

		role := c.Params("role")
		var req userModel.SignUpReq

		if role != "student" {
			req.GroupID = 0
		}

		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		userID, err := h.service.CreateUser(ctx, req, role)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully created user")
		return c.Status(fiber.StatusOK).JSON(userID)
	}
}

// Login handler method for login.
func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[UserHandler]").Start(c.Context(), "[Login]")
		defer span.End()

		role := c.Params("role")

		var loginReq userModel.LoginReq
		if err := reqvalidator.ReadRequest(c, &loginReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		userWithToken, err := h.service.Login(ctx, loginReq, role)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		utils.SetAuthCookies(c, h.cfg, userWithToken.AccessToken, userWithToken.RefreshToken)

		span.SetStatus(codes.Ok, "login successfully completed")
		return c.Status(fiber.StatusOK).JSON(userWithToken)
	}
}
