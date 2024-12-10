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
	FindByIDs(ctx context.Context, ids uuid.UUIDs) ([]*entity.User, error)
}

type RecommendationLogRepository interface {
	InsertLogs(ctx context.Context, logs []*entity.RecommendationLog) error
	FindTodayLogsByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.RecommendationLog, error)
}
