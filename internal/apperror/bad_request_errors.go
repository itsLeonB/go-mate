package apperror

func LoginError() *AppError {
	return BadRequestError("Wrong email/password", nil)
}
