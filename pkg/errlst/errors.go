package errlst

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Package errlst for debugging and handling
// errors in any layer of project!!

// errors are
var (
	// ErrBadRequest is
	ErrBadRequest = errors.New("bad request")
	// ErrBadQueryParams is
	ErrBadQueryParams = errors.New("bad query params")
	// ErrTypeAssertInTransaction is
	ErrTypeAssertInTransaction = errors.New("error in type assertion to connection.DBOps")
	// ErrBeginTransaction is
	ErrBeginTransaction = errors.New("cannot start transaction")
	// ErrUnauthorized is
	ErrUnauthorized = errors.New("unauthorized")
	// ErrNotFound is
	ErrNotFound = errors.New("not found")
	// ErrConflict is
	ErrConflict = errors.New("conflict occured")
	// ErrForbidden is
	ErrForbidden = errors.New("forbidden")
	// ErrFieldValidation is
	ErrFieldValidation = errors.New("field validation error")
	// ErrNoSuchUser is
	ErrNoSuchUser = errors.New("no such user")
	// ErrInternalServer is
	ErrInternalServer = errors.New("internal server error")
	// ErrTransactionFailed is
	ErrTransactionFailed = errors.New("failed to perform transaction")
	// ErrInvalidJWTToken is
	ErrInvalidJWTToken = errors.New("invalid JWT Token")
	// ErrTokenExpired is
	ErrTokenExpired = errors.New("token is expired")
	// ErrInvalidJWTMethod is
	ErrInvalidJWTMethod = errors.New("invalid jwt token method")
	// ErrInvalidJWTClaims is
	ErrInvalidJWTClaims = errors.New("invalid JWT Claims")

	// ErrRange is
	ErrRange = errors.New("value out of range")
	// ErrSyntax is
	ErrSyntax = errors.New("invalid syntax")
)

var _ RestErr = (*RestError)(nil)

// RestErr interface needs capturing errors
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	// AppearedAt() time.Time
}

// RestError struct to implement the RestErr interface
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrMessage string      `json:"err_msg,omitempty"`
	ErrCause   interface{} `json:"err_cause,omitempty"`
}

// ---------- IMPLEMENTING RestErr methods -------------
// Status is
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes is
func (e RestError) Causes() interface{} {
	return e.ErrCause
}

// ErrorMessage is
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - err_msg: %s - causes: %v - appearedAt: %v",
		e.ErrStatus, e.ErrMessage, e.ErrCause, time.Now())
}

// ------------------- FACTORY FUNCTIONS FOR CREATING ERRORS ---------------------
// NewBadRequestError creates a new 400 bad request error
func NewBadRequestError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCause:   cause,
	}
}

// NewNotFoundError creates a new 404 not found error
func NewNotFoundError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: ErrNotFound.Error(),
		ErrCause:   cause,
	}
}

// NewBadQueryParamsError creates a 403 bad query params error
func NewBadQueryParamsError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadQueryParams.Error(),
		ErrCause:   cause,
	}
}

// NewUnauthorizedError creates a 401 unauthorized error
func NewUnauthorizedError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrUnauthorized.Error(),
		ErrCause:   cause,
	}
}

// NewInternalServerError creates a 500 internal server error
func NewInternalServerError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServer.Error(),
		ErrCause:   cause,
	}
}

// NewConflictError creates a new 409 conflict error
func NewConflictError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrMessage: ErrConflict.Error(),
		ErrCause:   cause,
	}
}

// NewForbiddenError creates a new 403 forbidden error
func NewForbiddenError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: ErrForbidden.Error(),
		ErrCause:   cause,
	}
}
