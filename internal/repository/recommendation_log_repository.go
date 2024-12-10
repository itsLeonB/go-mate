package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/util"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type recommendationLogRepositoryGorm struct {
	db *gorm.DB
}

func NewRecommendationLogRepositoryGorm(db *gorm.DB) RecommendationLogRepository {
	return &recommendationLogRepositoryGorm{db}
}

func (rlrg *recommendationLogRepositoryGorm) InsertLogs(ctx context.Context, logs []*entity.RecommendationLog) error {
	err := rlrg.db.Create(&logs).Error
	if err != nil {
		return eris.Wrap(err, "error while inserting logs")
	}

	return nil
}

func (rlrg *recommendationLogRepositoryGorm) FindTodayLogsByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.RecommendationLog, error) {
	var logs []*entity.RecommendationLog
	now := time.Now()

	err := rlrg.db.Where("user_id = ? AND created_at >= ? AND created_at <= ?",
		userID, util.StartOfDay(now), util.EndOfDay(now),
	).Find(&logs).Error
	if err != nil {
		return nil, eris.Wrap(err, "error while finding logs by user id")
	}

	return logs, nil
}
