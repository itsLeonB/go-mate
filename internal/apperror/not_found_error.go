package apperror

import (
	"fmt"

	"github.com/google/uuid"
)

func UserNotFoundError(id uuid.UUID) *AppError {
	return NotFoundError(fmt.Sprintf("user with id %s is not found", id.String()))
}

func LogNotFoundError(recommendedUserID uuid.UUID) *AppError {
	return NotFoundError(fmt.Sprintf(
		"current user has not viewed user with id: %s yet",
		recommendedUserID.String(),
	))
}
