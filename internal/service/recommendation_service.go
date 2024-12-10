package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/mapper"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/repository"
	"github.com/rotisserie/eris"
)

type recommendationServiceNaive struct {
	userRepository              repository.UserRepository
	scoringService              ScoringService
	recommendationLogRepository repository.RecommendationLogRepository
}

func NewRecommendationServiceNaive(
	userRepository repository.UserRepository,
	scoringService ScoringService,
	recommendationLogRepository repository.RecommendationLogRepository,
) RecommendationService {
	return &recommendationServiceNaive{
		userRepository:              userRepository,
		scoringService:              scoringService,
		recommendationLogRepository: recommendationLogRepository,
	}
}

func (rsn *recommendationServiceNaive) GetUserRecommendations(ctx context.Context) ([]*model.UserResponse, error) {
	userID, err := uuid.Parse(ctx.Value(appconstant.ContextUserID).(string))
	if err != nil {
		return nil, eris.Wrap(err, "error while parsing user id")
	}

	todayLogs, err := rsn.recommendationLogRepository.FindTodayLogsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(todayLogs) > 0 {
		var availableToView uuid.UUIDs
		for _, log := range todayLogs {
			if log.Status == appconstant.LogStatusPending {
				availableToView = append(availableToView, log.RecommendedUserID)
			}
		}
		if len(availableToView) == 0 {
			return []*model.UserResponse{}, nil
		}

		todayRecommendedUsers, err := rsn.userRepository.FindByIDs(ctx, availableToView)
		if err != nil {
			return nil, err
		}

		return mapper.MapUsersToResponses(todayRecommendedUsers), nil
	}

	users, err := rsn.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	recommendedUsers := rsn.scoringService.ScoreAndSortUsers(ctx, users)
	recommendationLogs := mapper.MapUsersToRecommendationLogs(userID, recommendedUsers)

	err = rsn.recommendationLogRepository.InsertLogs(ctx, recommendationLogs)
	if err != nil {
		return nil, err
	}

	return mapper.MapUsersToResponses(recommendedUsers), nil
}
