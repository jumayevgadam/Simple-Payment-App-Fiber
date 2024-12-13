package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

var _ times.Handlers = (*TimeHandler)(nil)

type TimeHandler struct {
	service services.DataService
}

func NewTimeHandler(service services.DataService) *TimeHandler {
	return &TimeHandler{service: service}
}

func (t *TimeHandler) AddTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req timeModel.Request

		err := reqvalidator.ReadRequest(c, &req)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		timeID, err := t.service.TimeService().AddTime(c.Context(), req)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(timeID)
	}
}
