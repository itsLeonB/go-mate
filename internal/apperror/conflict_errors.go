package apperror

import (
	"fmt"
	"time"
)

func DuplicateEmailError(email string) *AppError {
	return ConflictError(fmt.Sprintf("email %s already registered", email), nil)
}

func LogAlreadyUpdatedError(userID string, recommendedUserID string) *AppError {
	return ConflictError(fmt.Sprintf(
		"user with id: %s has already interacted with user with id: %s for today",
		userID, recommendedUserID,
	), nil)
}

func SubscriptionAlreadyExistsError(model string, expiredAt time.Time) *AppError {
	return ConflictError(fmt.Sprintf(
		"user already has %s subscription, and will expired at: %s, please wait after expiry",
		model, expiredAt.String(),
	), nil)
}
