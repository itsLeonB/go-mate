package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByIDs(ctx context.Context, ids uuid.UUIDs) ([]*entity.User, error)
}

type RecommendationLogRepository interface {
	InsertLogs(ctx context.Context, logs []*entity.RecommendationLog) error
	FindTodayLogsByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.RecommendationLog, error)
	FindTodayLogByUserIDAndRecommendedUserID(ctx context.Context, userID uuid.UUID, recommendedUserID uuid.UUID) (*entity.RecommendationLog, error)
	Update(ctx context.Context, log *entity.RecommendationLog) error
}

type SubscriptionRepository interface {
	Insert(ctx context.Context, subscription *entity.UserSubscription) error
}
