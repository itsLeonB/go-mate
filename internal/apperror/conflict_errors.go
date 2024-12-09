package apperror

import "fmt"

func DuplicateEmailError(email string) *AppError {
	return ConflictError(fmt.Sprintf("email %s already registered", email), nil)
}
