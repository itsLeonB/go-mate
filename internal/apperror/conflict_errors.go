package apperror

import "fmt"

func DuplicateEmailError(email string) *AppError {
	return ConflictError(fmt.Sprintf("email %s already registered", email), nil)
}

func LogAlreadyUpdatedError(userID string, recommendedUserID string) *AppError {
	return ConflictError(fmt.Sprintf(
		"user with id: %s has already interacted with user with id: %s for today",
		userID, recommendedUserID,
	), nil)
}
