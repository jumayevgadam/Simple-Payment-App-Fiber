package reqvalidator

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ReadRequest body and validate
func ReadRequest(ctx *fiber.Ctx, request interface{}) error {
	if err := ctx.BodyParser(request); err != nil {
		return fmt.Errorf("error in reading request: %w", err)
	}

	if err := validate.StructCtx(ctx.Context(), request); err != nil {
		return fmt.Errorf("error in validating request: %w", err)
	}

	return nil
}
