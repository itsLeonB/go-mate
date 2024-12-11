package model

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
