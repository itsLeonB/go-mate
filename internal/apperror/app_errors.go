package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	HttpStatusCode int    `json:"-"`
	Type           string `json:"type"`
	Message        string `json:"message"`
	Details        any    `json:"details,omitempty"`
	Err            error  `json:"-"`
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s. %s", ae.HttpStatusCode, ae.Type, ae.Message, ae.Details)
}

func ConflictError(details any, err error) *AppError {
	return &AppError{
		HttpStatusCode: http.StatusConflict,
		Type:           "ConflictError",
		Message:        "Conflict on existing resource",
		Details:        details,
		Err:            err,
	}
}

func BadRequestError(details any, err error) *AppError {
	return &AppError{
		HttpStatusCode: http.StatusBadRequest,
		Type:           "BadRequestError",
		Message:        "Bad request",
		Details:        details,
		Err:            err,
	}
}

func InternalServerError() *AppError {
	return &AppError{
		HttpStatusCode: http.StatusInternalServerError,
		Type:           "InternalServerError",
		Message:        "Unexpected error occurred",
	}
}

func UnauthorizedError(details any) *AppError {
	return &AppError{
		HttpStatusCode: http.StatusUnauthorized,
		Type:           "UnauthorizedError",
		Message:        "Unauthorized",
		Details:        details,
	}
}
