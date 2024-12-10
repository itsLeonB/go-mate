package model

import "github.com/google/uuid"

type LogActionRequest struct {
	RecommendedUserID uuid.UUID `json:"recommended_user_id" binding:"required,uuid"`
	Action            string    `json:"action" binding:"required,oneof=liked passed"`
}

type RecommendationLogResponse struct {
	ID                string `json:"id"`
	UserID            string `json:"user_id"`
	RecommendedUserID string `json:"recommended_user_id"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
