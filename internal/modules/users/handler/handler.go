package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
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

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginRequest userModel.LoginRequest

		err := reqvalidator.ReadRequest(c, &loginRequest)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		loginResponse, err := h.service.UserService().Login(c.Context(), loginRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(loginResponse)
	}
}
