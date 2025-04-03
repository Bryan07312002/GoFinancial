package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FieldError represents a single field error
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ServiceError represents the error response structure
type ServiceError struct {
	StatusCode int          `json:"-"`
	Message    string       `json:"message,omitempty"`
	Errors     []FieldError `json:"errors,omitempty"`
	Details    any          `json:"details,omitempty"`
}

// New creates a new ServiceError with a status code and message
func New(statusCode int, message string) *ServiceError {
	return &ServiceError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// AddFieldError adds a field-specific error
func (e *ServiceError) AddFieldError(field, message string) *ServiceError {
	e.Errors = append(e.Errors, FieldError{
		Field:   field,
		Message: message,
	})
	return e
}

// WithDetails adds additional details to the error
func (e *ServiceError) WithDetails(details any) *ServiceError {
	e.Details = details
	return e
}

// Error implements the error interface
func (e *ServiceError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("%s: %d field errors", e.Message, len(e.Errors))
	}

	return e.Message
}

// ToJSON converts the error to JSON
func (e *ServiceError) ToJSON() []byte {
	jsonData, _ := json.Marshal(e)
	return jsonData
}

// WriteJSON writes the error as JSON to the http response
func (e *ServiceError) WriteJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	w.Write(e.ToJSON())
}

// Common field error creators
func RequiredField(field string) FieldError {
	return FieldError{
		Field:   field,
		Message: fmt.Sprintf("%s is required", field),
	}
}

func InvalidFormat(field string) FieldError {
	return FieldError{
		Field:   field,
		Message: fmt.Sprintf("%s has invalid format", field),
	}
}

func TooShort(field string, min int) FieldError {
	return FieldError{
		Field:   field,
		Message: fmt.Sprintf("%s must be at least %d characters", field, min),
	}
}

func TooLong(field string, max int) FieldError {
	return FieldError{
		Field:   field,
		Message: fmt.Sprintf("%s must be at most %d characters", field, max),
	}
}

const (
	// Common error messages
	ErrInternalServer     = "internal server error"
	ErrBadRequest         = "invalid request"
	ErrUnauthorized       = "unauthorized"
	ErrForbidden          = "forbidden"
	ErrNotFound           = "not found"
	ErrConflict           = "conflict"
	ErrValidationFailed   = "validation failed"
	ErrInvalidCredentials = "invalid credentials"
	ErrExpiredToken       = "expired token"
	ErrInvalidToken       = "invalid token"
)

// Common error constructors
func InternalServerError() *ServiceError {
	return New(http.StatusInternalServerError, ErrInternalServer)
}

func BadRequestError() *ServiceError {
	return New(http.StatusBadRequest, ErrBadRequest)
}

func UnauthorizedError() *ServiceError {
	return New(http.StatusUnauthorized, ErrUnauthorized)
}

func ForbiddenError() *ServiceError {
	return New(http.StatusForbidden, ErrForbidden)
}

func NotFoundError() *ServiceError {
	return New(http.StatusNotFound, ErrNotFound)
}

func ConflictError() *ServiceError {
	return New(http.StatusConflict, ErrConflict)
}

func ValidationError() *ServiceError {
	return New(http.StatusBadRequest, ErrValidationFailed)
}
