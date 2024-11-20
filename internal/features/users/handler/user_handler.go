package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	userOps "github.com/jumayevgadaym/tsu-toleg/internal/features/users"
	userModel "github.com/jumayevgadaym/tsu-toleg/internal/models/user"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadaym/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadaym/tsu-toleg/pkg/utils"
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
func (h *UserHandler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[UserHandler]").Start(c.Context(), "[CreateUser]")
		defer span.End()

		var req userModel.SignUpReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		userID, err := h.service.CreateUser(ctx, req)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully created user")
		return c.Status(fiber.StatusOK).JSON(userID)
	}
}

// Login handler method for login.
func (h *UserHandler) Login(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[UserHandler]").Start(c.Context(), "[Login]")
		defer span.End()

		var loginReq userModel.LoginReq
		if err := reqvalidator.ReadRequest(c, &loginReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		userWithToken, err := h.service.Login(ctx, loginReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		utils.SetAuthCookies(c, h.cfg, userWithToken.AccessToken, userWithToken.RefreshToken)

		span.SetStatus(codes.Ok, "login successfully completed")
		return c.Status(fiber.StatusOK).JSON(userWithToken)
	}
}
