package mapper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/entity"
)

func MapUserToRecommendationLog(userID uuid.UUID, user *entity.User) *entity.RecommendationLog {
	return &entity.RecommendationLog{
		UserID:            userID,
		RecommendedUserID: user.ID,
		Status:            "pending",
	}
}

func MapUsersToRecommendationLogs(userID uuid.UUID, users []*entity.User) []*entity.RecommendationLog {
	logs := make([]*entity.RecommendationLog, len(users))
	for i := 0; i < len(users); i++ {
		logs[i] = MapUserToRecommendationLog(userID, users[i])
	}

	return logs
}
