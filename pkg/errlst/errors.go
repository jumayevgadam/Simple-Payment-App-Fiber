package errlst

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Package errlst provides a centralized place for defining errors used in various layers of the project.

// Common errors encountered throughout the application:

var (
	// ErrBadRequest is used when a client sends an invalid request.
	ErrBadRequest = errors.New("bad request")

	// ErrBadQueryParams is used when the query parameters provided by the client are invalid.
	ErrBadQueryParams = errors.New("bad query params")

	// ErrTypeAssertInTransaction is used when there is an error in type assertion during database transaction operations.
	ErrTypeAssertInTransaction = errors.New("error in type assertion to connection.DBOps")

	// ErrBeginTransaction is returned when the application fails to start a new transaction.
	ErrBeginTransaction = errors.New("cannot start transaction")

	// ErrUnauthorized is returned when a user attempts to access a resource without proper authorization.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrNotFound is returned when the requested resource could not be found in the system.
	ErrNotFound = errors.New("not found")

	// ErrConflict is used when there is a conflict with the current state of the system.
	ErrConflict = errors.New("conflict occurred")

	// ErrForbidden is used when a user tries to access a resource they do not have permission to access.
	ErrForbidden = errors.New("forbidden")

	// ErrFieldValidation is used when there is a validation error on input fields (e.g., invalid or missing data in a request).
	ErrFieldValidation = errors.New("field validation error")

	// ErrNoSuchUser is returned when an attempt is made to find a user that does not exist in the system.
	ErrNoSuchUser = errors.New("no such user")

	// ErrNoSuchRole is returned when an attempt is made to find roles associated with a permission that does not exist.
	ErrNoSuchRole = errors.New("no roles found for this permission")

	// ErrInternalServer is returned when an unexpected internal server error occurs.
	ErrInternalServer = errors.New("internal server error")

	// ErrTransactionFailed is used when a database transaction fails, often due to unexpected errors during operations.
	ErrTransactionFailed = errors.New("failed to perform transaction")

	// ErrInvalidJWTToken is returned when a JWT token is invalid (e.g., expired, malformed, or incorrectly signed).
	ErrInvalidJWTToken = errors.New("invalid JWT Token")

	// ErrTokenExpired is used when a JWT token has expired and is no longer valid.
	ErrTokenExpired = errors.New("token is expired")

	// ErrInvalidJWTMethod is returned when an unsupported or invalid method is used with a JWT token.
	ErrInvalidJWTMethod = errors.New("invalid jwt token method")

	// ErrInvalidJWTClaims is used when the claims within a JWT token are invalid or do not meet expected criteria.
	ErrInvalidJWTClaims = errors.New("invalid JWT Claims")

	// ErrRange is returned when a value is outside the acceptable range (e.g., too high or too low).
	ErrRange = errors.New("value out of range")

	// ErrSyntax is returned when there is a syntax error in the input or the request.
	ErrSyntax = errors.New("invalid syntax")
)

var _ RestErr = (*RestError)(nil) //nolint:errcheck

// RestErr interface needs capturing errors.
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	// AppearedAt() time.Time
}

// RestError struct to implement the RestErr interface.
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrMessage string      `json:"err_msg,omitempty"`
	ErrCause   interface{} `json:"err_cause,omitempty"`
}

// ---------- IMPLEMENTING RestErr methods -------------.

// Status is.
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes is.
func (e RestError) Causes() interface{} {
	return e.ErrCause
}

// ErrorMessage is.
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - err_msg: %s - causes: %v - appearedAt: %v",
		e.ErrStatus, e.ErrMessage, e.ErrCause, time.Now())
}

// ------------------- FACTORY FUNCTIONS FOR CREATING ERRORS ---------------------.

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
