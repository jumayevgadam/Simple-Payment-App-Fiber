package errlst

import (
	"errors"
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

	// ErrNoUploadedFile is used when student does not upload payment.
	ErrNoUploadedFile = errors.New("there is no uploaded file associated with the given key")

	// ErrFileSize is returned when student uploaded file bigger than 5MB.
	ErrFileSize = errors.New("file size exceeds from limit 5MB")

	// ------------------ NOT FOUND COLLECTIOIN OF ERRORS ----------------------------------------.

	// ErrNotFound is returned when the requested resource could not be found in the system.
	ErrNotFound = errors.New("not found")

	// ErrFileNotFound is returned when the uploaded file not found.
	ErrFileNotFound = errors.New("file not found")

	// ErrStudentNotFound is returned when student not found an identified id.
	ErrStudentNotFound = errors.New("student not found")

	// ErrAdminNotFound is returned when admin not found an identified id.
	ErrAdminNotFound = errors.New("admin not found")

	// ErrPaymentNotFound is returned whene there is not a payment in given id.
	ErrPaymentNotFound = errors.New("no payment found with the given paymentID")

	// ErrNoSuchRole is returned when an attempt is made to find roles associated with a permission that does not exist.
	ErrNoSuchRole = errors.New("no roles found for this permission")

	// ErrGroupNotFound is returned when group not found in that identified id.
	ErrGroupNotFound = errors.New("group not found an identified id")

	// ErrFacultyNotFound is returned when faculty not found in that identified id.
	ErrFacultyNotFound = errors.New("faculty not found an identified id")

	// ErrActiveYearNotFound is returned when active year not setted in database.
	ErrActiveYearNotFound = errors.New("active year not found for performing this action")

	// -------------------- SPECIFIC ERRORS FOR PAYMENT MODULE -----------------------------------.

	// ErrPaymentPerformedForThisYear is used when student also perform payment after full payment.
	ErrPaymentPerform = errors.New("")

	// ErrDidNotPerformPayment is used when student did not perform that academic year yet.
	ErrDidNotPerformPayment = errors.New("birinji semestriň tölegini ýada doly tölegi ýerine ýetiriň")

	// ErrFullPayment is used when student perform full payment with wrong details.
	ErrFullPayment = errors.New("doly töleg üçin nädogry balans girizildi")

	// ErrFirstSemesterPayment is used when student perform wrong payment for first semester.
	ErrFirstSemesterPayment = errors.New("birinji semestr tölegi üçin dogry balans girizilmedi")

	// ErrDidNotPerformFullPayment is used when student perform first semester payment, then does not perform full payment.
	ErrDidNotPerformFullPayment = errors.New("birinji semestr tölendi, ikinji semesteri töläň")

	// ErrSecondSemesterPayment is used when student perform wrong payment for second semester.
	ErrSecondSemesterPayment = errors.New("ikinji semesteriň tölegi üçin dogry balans girizilmedi")

	// ErrInPaymentType is used when student perform wrong payment type.
	ErrInPaymentType = errors.New("töleg etmek üçin semestr saýlanmady")

	// ErrMisMatchedStudentID is used when payments.student_id is not equal to studentID.
	ErrMisMatchedStudentID = errors.New("studentID mismatched with payments.student_id")

	// ------------------ SPECIFIC ERRORS FOR GROUPS MODULE --------------------------------------.
)
