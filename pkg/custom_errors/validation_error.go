package custom_errors

import (
	"errors"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// FieldError represents structured validation errors
type FieldError struct {
	Message string   `json:"message"`
	Types   []string `json:"types"`
}

// ParseValidationTypes extracts validation rule names from error messages
func ParseValidationTypes(err error) []string {
	var types []string
	if err == nil {
		return types
	}

	errorMsg := err.Error()
	if validation.ErrRequired.Error() == errorMsg {
		types = append(types, "required")
	}
	if validation.ErrLengthOutOfRange.Error() == errorMsg {
		types = append(types, "length")
	}
	if validation.ErrLengthTooShort.Error() == errorMsg {
		types = append(types, "too_short")
	}
	if validation.ErrLengthTooLong.Error() == errorMsg {
		types = append(types, "too_long")
	}
	if is.ErrURL.Error() == errorMsg {
		types = append(types, "url")
	}

	return types
}

// MapValidationErrors converts Ozzo validation errors into structured format
func MapValidationErrors(err error) map[string]FieldError {
	validationErrors := make(map[string]FieldError)

	if err == nil {
		return nil
	}

	var errs validation.Errors
	if errors.As(err, &errs) {
		for field, fieldErr := range errs {
			if fieldErr != nil {
				validationErrors[field] = FieldError{
					Message: fieldErr.Error(),
					Types:   ParseValidationTypes(fieldErr),
				}
			}
		}
	}

	return validationErrors
}
