package errlst

import (
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

// RestErr interface needs capturing errors.
type RestErr interface {
	Status() int
	Error() string
}

// RestError struct to implement the RestErr interface.
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrKind    string      `json:"err_kind,omitempty"`
	ErrMessage interface{} `json:"err_msg,omitempty"`
}

// ---------- IMPLEMENTING RestErr methods -------------.

// Status is.
func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Error() string {
	return fmt.Sprintf("%v", e.ErrMessage)
}

// NewBadRequestError creates a new 400 bad request error.
func NewBadRequestError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrKind:    ErrBadRequest.Error(),
		ErrMessage: err,
	}
}

// NewNotFoundError creates a new 404 not found error.
func NewNotFoundError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrKind:    ErrNotFound.Error(),
		ErrMessage: err,
	}
}

// NewBadQueryParamsError creates a 403 bad query params error.
func NewBadQueryParamsError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrKind:    ErrBadQueryParams.Error(),
		ErrMessage: err,
	}
}

// NewUnauthorizedError creates a 401 unauthorized error.
func NewUnauthorizedError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrKind:    ErrUnauthorized.Error(),
		ErrMessage: err,
	}
}

// NewInternalServerError creates a 500 internal server error.
func NewInternalServerError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrKind:    ErrInternalServer.Error(),
		ErrMessage: err,
	}
}

// NewConflictError creates a new 409 conflict error.
func NewConflictError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrKind:    ErrConflict.Error(),
		ErrMessage: err,
	}
}

// NewForbiddenError creates a new 403 forbidden error.
func NewForbiddenError(err any) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrKind:    ErrForbidden.Error(),
		ErrMessage: err,
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
	case errors.Is(err, pgx.ErrNoRows):
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
