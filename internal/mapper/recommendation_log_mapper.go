package mapper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/model"
)

func NewRecommendationLog(userID uuid.UUID, user *entity.User) *entity.RecommendationLog {
	return &entity.RecommendationLog{
		UserID:            userID,
		RecommendedUserID: user.ID,
		Status:            appconstant.LogStatusViewed,
	}
}

func NewRecommendationLogs(userID uuid.UUID, users []*entity.User) []*entity.RecommendationLog {
	logs := make([]*entity.RecommendationLog, len(users))
	for i := 0; i < len(users); i++ {
		logs[i] = NewRecommendationLog(userID, users[i])
	}

	return logs
}

func MapRecommendationLogToResponse(log *entity.RecommendationLog) *model.RecommendationLogResponse {
	return &model.RecommendationLogResponse{
		ID:                log.ID.String(),
		UserID:            log.UserID.String(),
		RecommendedUserID: log.RecommendedUserID.String(),
		Status:            log.Status,
		CreatedAt:         log.CreatedAt.String(),
		UpdatedAt:         log.UpdatedAt.String(),
	}
}
