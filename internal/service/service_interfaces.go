package service

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, request *model.RegisterRequest) error
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	ValidateUser(ctx context.Context) (*entity.User, error)
}

type RecommendationService interface {
	GetUserRecommendations(ctx context.Context) ([]*model.UserResponse, error)
	LogAction(ctx context.Context, request *model.LogActionRequest) (*model.RecommendationLogResponse, error)
}

type ScoringService interface {
	ScoreAndSortUsers(ctx context.Context, users []*entity.User, isExtraRecommendation bool) ([]*entity.User, error)
}

type SubscriptionService interface {
	AddSubscription(ctx context.Context, request *model.NewSubscriptionRequest) (*model.UserSubscriptionResponse, error)
}
