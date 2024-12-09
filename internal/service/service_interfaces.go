package service

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, request *model.RegisterRequest) error
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}

type RecommendationService interface {
	GetUserRecommendations(ctx context.Context) ([]*model.UserResponse, error)
}

type ScoringService interface {
	ScoreAndSortUsers(ctx context.Context, users []*entity.User) []*entity.User
}
