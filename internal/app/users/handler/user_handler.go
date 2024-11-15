package handler

import (
	"github.com/gofiber/fiber/v2"
	userOps "github.com/jumayevgadaym/tsu-toleg/internal/app/users"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	userModel "github.com/jumayevgadaym/tsu-toleg/internal/models/user"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadaym/tsu-toleg/pkg/reqvalidator"
	"github.com/jumayevgadaym/tsu-toleg/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	_ userOps.Handler = (*UserHandler)(nil)
)

// UserHandler is
type UserHandler struct {
	cfg     *config.Config
	service userOps.Service
}

// NewUserHandler is
func NewUserHandler(cfg *config.Config, service userOps.Service) *UserHandler {
	return &UserHandler{cfg: cfg, service: service}
}

// CreateUser handler is
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

// Login handler is
func (h *UserHandler) Login() fiber.Handler {
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

		// SetTOCookie tokens
		utils.SetAuthCookies(&h.cfg.JWT, userWithToken.AccessToken, userWithToken.RefreshToken)

		span.SetStatus(codes.Ok, "login successfully completed")
		return c.Status(fiber.StatusOK).JSON(userWithToken)
	}
}
