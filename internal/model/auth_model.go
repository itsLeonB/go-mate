package model

type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
		Type:  "Bearer",
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func NewRegisterResponse() *RegisterResponse {
	return &RegisterResponse{
		Message: "register success, please login",
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
