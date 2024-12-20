package errlst

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// NewBadRequestError creates a new 400 bad request error.
func NewBadRequestError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCause:   cause,
	}
}

// NewNotFoundError creates a new 404 not found error.
func NewNotFoundError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: ErrNotFound.Error(),
		ErrCause:   cause,
	}
}

// NewBadQueryParamsError creates a 403 bad query params error.
func NewBadQueryParamsError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadQueryParams.Error(),
		ErrCause:   cause,
	}
}

// NewUnauthorizedError creates a 401 unauthorized error.
func NewUnauthorizedError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrUnauthorized.Error(),
		ErrCause:   cause,
	}
}

// NewInternalServerError creates a 500 internal server error.
func NewInternalServerError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServer.Error(),
		ErrCause:   cause,
	}
}

// NewConflictError creates a new 409 conflict error.
func NewConflictError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrMessage: ErrConflict.Error(),
		ErrCause:   cause,
	}
}

// NewForbiddenError creates a new 403 forbidden error.
func NewForbiddenError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: ErrForbidden.Error(),
		ErrCause:   cause,
	}
}

// ParseSQLErrors returns corresponding RestErr.
func ParseSQLErrors(err error) RestErr {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// CLASS 22
		case "22001": // numeric value out of range
			return NewBadRequestError("Numeric" + ErrRange.Error() + pgErr.Message + "\n")

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + ErrConflict.Error() + pgErr.Message + "\n")
		case "23503": // Foreign key violation
			return NewBadRequestError("Foreign key violation: " + pgErr.Message + "\n")
		case "23502": // Not null violation
			return NewBadRequestError("Not null violation: " + pgErr.Message + "\n")

		// CLASS 40
		case "40001": // serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message + "\n")
		// CLASS 42
		case "42601": // syntax error
			return NewBadRequestError("Syntax error in sql statements: " + pgErr.Message + "\n")
		}
	}

	if strings.Contains(err.Error(), "scany") {
		return NewBadRequestError(err.Error() + "\n")
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error() + "\n")
	}

	return NewBadRequestError(err.Error() + "\n")
}

// ParseValidatorError parses validation errors and returns corresponding RestErr.
func ParseValidatorError(err error) RestErr {
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		return NewBadRequestError(err.Error())
	}

	// Pre-allocate the errorMessages slice
	errorMessages := make([]string, 0, len(validationErrs))

	for _, fieldErr := range validationErrs {
		// For each validation error, create a message.
		errorMessage := fmt.Sprintf("Field validation for %s, failed on the %s tag",
			fieldErr.Field(), fieldErr.Tag())
		// append each error to error messages
		errorMessages = append(errorMessages, errorMessage)
	}

	// Combine all messages into one string and return a BadRequestError.
	return NewBadRequestError(strings.Join(errorMessages, ", "))
}

// ParseErrors parses common errors (like SQL errors) into the RestErr.
func ParseErrors(err error) RestErr {
	switch {
	// pgx specific errors
	case errors.Is(err, sql.ErrNoRows):
		return NewNotFoundError(err.Error())
	case errors.Is(err, pgx.ErrTooManyRows):
		return NewConflictError(err.Error())
	case strings.Contains(err.Error(), "not found"):
		return NewNotFoundError(err.Error())

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
		return NewUnauthorizedError(ErrUnauthorized.Error() + err.Error() + "\n")

	default:
		// If already a RestErr, return as-is
		var restErr RestErr
		if errors.As(err, &restErr) {
			return restErr
		}

		return NewInternalServerError("internal server error: [ParseErrors]" + err.Error() + "\n")
	}
}

// Response returns ErrorResponse, for clean syntax I took function name Response.
// Because in every package i call this errlst package httpError, serviceErr, repoErr.
// Then easily call this httpError.Response(err), serviceErr.Response(err), repoErr.Response(err).
func Response(c *fiber.Ctx, err error) error {
	logger := logger.NewAPILogger(&config.Config{})
	logger.InitLogger()

	errStatus, errResponse := ParseErrors(err).Status(), ParseErrors(err)
	logger.Error("HTTP Error Response: ", err, ", URL:", c.OriginalURL())

	return c.Status(errStatus).JSON(errResponse)
}
