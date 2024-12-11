package model

type NewSubscriptionRequest struct {
	Model string `json:"model" binding:"required,oneof=extra_recommendations extra_appearance"`
	Plan  string `json:"plan" binding:"required,oneof=monthly yearly"`
}

type UserSubscriptionResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Model     string `json:"model"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ExpiredAt string `json:"expiredAt"`
}
