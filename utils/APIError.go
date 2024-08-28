package utils

import (
	"fmt"
	"net/http"
)

// APIError represents an error that occurred during an API request
type APIError struct {
	StatusCode int
	Code       string
	Message    string
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("API error: status %d, code %s, message: %s", e.StatusCode, e.Code, e.Message)
}

// NewAPIError creates a new APIError
func NewAPIError(statusCode int, code, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

// Helper functions for common HTTP errors
func NewNotFoundError(message string) *APIError {
	return NewAPIError(http.StatusNotFound, "NOT_FOUND", message)
}

func NewBadRequestError(message string) *APIError {
	return NewAPIError(http.StatusBadRequest, "BAD_REQUEST", message)
}

func NewInternalServerError(message string) *APIError {
	return NewAPIError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message)
}
