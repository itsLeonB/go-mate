package apperror

func MissingTokenError() *AppError {
	return UnauthorizedError("missing token")
}

func InvalidTokenError() *AppError {
	return UnauthorizedError("invalid token")
}
