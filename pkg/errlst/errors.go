package errlst

import (
	"errors"
	"fmt"
)

// Package errlst provides a centralized place for defining errors used in various layers of the project.

// Common errors encountered throughout the application.
var (
	// ------------------ COMMON APP, SERVER ERRORS ----------------------------------------------.

	// ErrConflict is used when there is a conflict with the current state of the system.
	ErrConflict = errors.New("conflict occurred")

	// ErrForbidden is used when a user tries to access a resource they do not have permission to access.
	ErrForbidden = errors.New("forbidden")

	// ErrInternalServer is returned when an unexpected internal server error occurs.
	ErrInternalServer = errors.New("internal server error")

	// ErrBadRequest is used when a client sends an invalid request.
	ErrBadRequest = errors.New("bad request")

	// ErrBadQueryParams is used when the query parameters provided by the client are invalid.
	ErrBadQueryParams = errors.New("bad query params")

	// ------------------ DATABASE AND TRANSACTION ERRORS ----------------------------------------.

	// ErrTypeAssertInTransaction is used when there is an error in type assertion during database transaction operations.
	ErrTypeAssertInTransaction = errors.New("error in type assertion to connection.DBOps")

	// ErrBeginTransaction is returned when the application fails to start a new transaction.
	ErrBeginTransaction = errors.New("cannot start transaction")

	// ErrTransactionFailed is used when a database transaction fails, often due to unexpected errors during operations.
	ErrTransactionFailed = errors.New("failed to perform transaction")

	// ------------------ JWT TOKEN AND AUTHORIZATION ERRORS -------------------------------------.

	// ErrAuthorizationHeaderNotProvided is used when authorization header not provided.
	ErrAuthorizationHeaderNotProvided = errors.New("authorization header is not provided")

	// ErrUnauthorized is returned when a user attempts to access a resource without proper authorization.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrInvalidJWTToken is returned when a JWT token is invalid (e.g., expired, malformed, or incorrectly signed).
	ErrInvalidJWTToken = errors.New("invalid JWT Token")

	// ErrTokenExpired is used when a JWT token has expired and is no longer valid.
	ErrTokenExpired = errors.New("token is expired")

	// ErrInvalidJWTMethod is returned when an unsupported or invalid method is used with a JWT token.
	ErrInvalidJWTMethod = errors.New("invalid jwt token method")

	// ErrInvalidJWTClaims is used when the claims within a JWT token are invalid or do not meet expected criteria.
	ErrInvalidJWTClaims = errors.New("invalid JWT Claims")

	// ------------------ FIELD VALIDATION ERRORS ------------------------------------------------.

	// ErrFieldValidation is used when there is a validation error on input fields (e.g., invalid or missing data in a request).
	ErrFieldValidation = errors.New("field validation error")

	// ErrRange is returned when a value is outside the acceptable range (e.g., too high or too low).
	ErrRange = errors.New("value out of range")

	// ErrSyntax is returned when there is a syntax error in the input or the request.
	ErrSyntax = errors.New("invalid syntax")

	// ------------------ NOT FOUND COLLECTIOIN OF ERRORS ----------------------------------------.

	// ErrNotFound is returned when the requested resource could not be found in the system.
	ErrNotFound = errors.New("not found")

	// ErrFileNotFound is returned when the uploaded file not found.
	ErrFileNotFound = errors.New("file not found")

	// ErrStudentNotFound is returned when student not found in that id.
	ErrStudentNotFound = errors.New("student not found")

	// ErrPaymentNotFound is returned whene there is not a payment in given id.
	ErrPaymentNotFound = errors.New("no payment found with the given id")

	// ErrNoSuchRole is returned when an attempt is made to find roles associated with a permission that does not exist.
	ErrNoSuchRole = errors.New("no roles found for this permission")

	// -------------------- SPECIFIC ERRORS FOR PAYMENT MODULE ---------------------------

	// ErrPaymentPerformedForThisYear is used when student also perform payment after full payment.
	ErrPaymentPerformedForThisYear = errors.New("you cannot perform full for this year, because this action has already been performed")
)

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

func (e RestError) Error() string {
	return fmt.Sprintf(
		"Error: {\n  Status: %d,\n  Message: \"%s\",\n  Cause: %v\n}",
		e.ErrStatus, e.ErrMessage, e.ErrCause,
	)
}
