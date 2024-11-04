package errlst

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ParseValidatorError parses validation errors and returns corresponding RestErr
func ParseValidatorError(err error) RestErr {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewBadRequestError(err.Error())
	}

	// Collect detailed validation error messages
	var errorMessages []string
	for _, fieldErr := range validationErrs {
		// For each validation error, create a message
		errorMessage := fmt.Sprintf("Field validation for %s, failed on the %s tag",
			fieldErr.Field(), fieldErr.Tag())
		// append each error to error messages
		errorMessages = append(errorMessages, errorMessage)
	}

	// Combine all messages into one string and return a BadRequestError
	return NewBadRequestError(strings.Join(errorMessages, ", "))
}
