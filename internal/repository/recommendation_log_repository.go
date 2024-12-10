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

func (rlrg *recommendationLogRepositoryGorm) FindTodayLogByUserIDAndRecommendedUserID(
	ctx context.Context,
	userID uuid.UUID,
	recommendedUserID uuid.UUID,
) (*entity.RecommendationLog, error) {
	var log entity.RecommendationLog
	now := time.Now()

	err := rlrg.db.Where("user_id = ? AND recommended_user_id = ? AND created_at >= ? AND created_at <= ?",
		userID, recommendedUserID, util.StartOfDay(now), util.EndOfDay(now),
	).First(&log).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error while finding log by recommended user id")
	}

	return &log, nil
}

func (rlrg *recommendationLogRepositoryGorm) Update(ctx context.Context, log *entity.RecommendationLog) error {
	err := rlrg.db.Save(log).Error
	if err != nil {
		return eris.Wrap(err, "error while updating log")
	}

	return nil
}
