package utils

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"tcfback/pkg/custom_errors"
)

// BaseResponse is a standardized API response format
type BaseResponse[T any] struct {
	RequestID string                 `json:"request_id,omitempty"`
	Message   string                 `json:"message"`
	Data      *T                     `json:"data,omitempty"`
	Errors    map[string]interface{} `json:"errors,omitempty"`
}

// SuccessResponse Standardized API Response
func SuccessResponse[T any](c echo.Context, statusCode int, message string, data *T) error {
	response := BaseResponse[T]{
		RequestID: generateRequestID(c),
		Message:   message,
		Data:      data,
		Errors:    nil,
	}
	return c.JSON(statusCode, response)
}

// ErrorResponse returns a standardized error response
func ErrorResponse(c echo.Context, statusCode int, message string, errors interface{}) error {
	// Convert validation errors to a generic map[string]interface{}
	var formattedErrors map[string]interface{}

	// if errors == nil
	if errors == nil {
		formattedErrors = nil
		return c.JSON(statusCode, BaseResponse[any]{
			RequestID: generateRequestID(c),
			Message:   message,
			Data:      nil,
			Errors:    nil,
		})
	}

	if errMap, ok := errors.(map[string]custom_errors.FieldError); ok {
		formattedErrors = make(map[string]interface{})
		for key, fieldErr := range errMap {
			formattedErrors[key] = fieldErr
		}
	} else {
		formattedErrors = errors.(map[string]interface{})
	}

	response := BaseResponse[any]{
		RequestID: generateRequestID(c),
		Message:   message,
		Data:      nil,
		Errors:    formattedErrors,
	}
	return c.JSON(statusCode, response)
}

// Generate a unique request ID if not provided
func generateRequestID(c echo.Context) string {
	requestID := c.Request().Header.Get("X-Request-ID")
	if requestID == "" {
		newUUID, _ := uuid.NewV7()
		requestID = newUUID.String()
	}
	return requestID
}
