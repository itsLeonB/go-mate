package mapper

import (
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
)

func MapUserSubscriptionToResponse(userSubscription *entity.UserSubscription) *model.UserSubscriptionResponse {
	return &model.UserSubscriptionResponse{
		ID:        userSubscription.ID.String(),
		UserID:    userSubscription.UserID.String(),
		Model:     userSubscription.Model,
		CreatedAt: userSubscription.CreatedAt.String(),
		UpdatedAt: userSubscription.UpdatedAt.String(),
		ExpiredAt: userSubscription.ExpiredAt.String(),
	}
}
