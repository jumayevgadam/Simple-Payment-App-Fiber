package errlst

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// ParseErrors parses common errors (like SQL errors) into the RestErr.
func ParseErrors(err error) RestErr {
	switch {
	// pgx specific errors
	case errors.Is(err, pgx.ErrNoRows):
		return NewNotFoundError(err.Error())
	case errors.Is(err, pgx.ErrTooManyRows):
		return NewConflictError(err.Error())

	// SQLSTATE error
	case strings.Contains(err.Error(), "SQLSTATE"):
		return ParseSQLErrors(err)

	// Handle strconv.Atoi errors
	case strings.Contains(err.Error(), ErrSyntax.Error()),
		strings.Contains(err.Error(), ErrRange.Error()):
		return NewBadRequestError(err.Error())

		// Handle Validation errors from go-validator/v10
	case errors.As(err, &validator.ValidationErrors{}):
		return ParseValidatorError(err)

	// Handle Token or Cookie errors
	case
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTToken.Error()),
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTClaims.Error()):
		return NewUnauthorizedError(ErrUnauthorized.Error() + err.Error())

	default:
		// If already a RestErr, return as-is
		var restErr RestErr
		if errors.As(err, &restErr) {
			return restErr
		}

		return NewInternalServerError("internal server error: [ParseErrors]" + err.Error())
	}
}

// Response returns ErrorResponse, for clean syntax I took function name Response.
// Because in every package i call this errlst package httpError, serviceErr, repoErr.
// Then easily call this httpError.Response(err), serviceErr.Response(err), repoErr.Response(err).
func Response(c *fiber.Ctx, err error) error {
	logger := logger.NewAPILogger(&config.Config{})
	logger.InitLogger()

	errStatus, errResponse := ParseErrors(err).Status(), ParseErrors(err)
	logger.Error("HTTP Error Response: ", err, c.Context().RemoteAddr())

	return c.Status(errStatus).JSON(errResponse)
}
