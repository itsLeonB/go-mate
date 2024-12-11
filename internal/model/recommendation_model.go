package model

import "github.com/google/uuid"

type LogActionRequest struct {
	RecommendedUserID uuid.UUID `json:"recommendedUserId" binding:"required,uuid"`
	Action            string    `json:"action" binding:"required,oneof=liked passed"`
}

type RecommendationLogResponse struct {
	ID                string `json:"id"`
	UserID            string `json:"userId"`
	RecommendedUserID string `json:"recommendedUserId"`
	Status            string `json:"status"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
