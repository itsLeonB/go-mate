package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/apperror"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/mapper"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/repository"
	"github.com/itsLeonB/go-mate/internal/util"
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
	user, err := rsn.validateUser(ctx)
	if err != nil {
		return nil, err
	}

	todayLogs, err := rsn.recommendationLogRepository.FindTodayLogsByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if len(todayLogs) > 0 {
		var availableToView uuid.UUIDs
		for _, log := range todayLogs {
			if log.Status == appconstant.LogStatusViewed {
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

	recommendedUsers, err := rsn.scoringService.ScoreAndSortUsers(ctx, users)
	if err != nil {
		return nil, err
	}
	if len(recommendedUsers) == 0 {
		return []*model.UserResponse{}, nil
	}

	recommendationLogs := mapper.NewRecommendationLogs(user.ID, recommendedUsers)

	err = rsn.recommendationLogRepository.InsertLogs(ctx, recommendationLogs)
	if err != nil {
		return nil, err
	}

	return mapper.MapUsersToResponses(recommendedUsers), nil
}

func (rsn *recommendationServiceNaive) LogAction(
	ctx context.Context,
	request *model.LogActionRequest,
) (*model.RecommendationLogResponse, error) {
	user, err := rsn.validateUser(ctx)
	if err != nil {
		return nil, err
	}

	log, err := rsn.recommendationLogRepository.FindTodayLogByUserIDAndRecommendedUserID(ctx, user.ID, request.RecommendedUserID)
	if err != nil {
		return nil, err
	}
	if log == nil {
		return nil, apperror.LogNotFoundError(request.RecommendedUserID)
	}
	if log.Status != appconstant.LogStatusViewed {
		return nil, apperror.LogAlreadyUpdatedError(user.ID.String(), request.RecommendedUserID.String())
	}

	log.Status = request.Action
	err = rsn.recommendationLogRepository.Update(ctx, log)
	if err != nil {
		return nil, err
	}

	return mapper.MapRecommendationLogToResponse(log), nil
}

func (rsn *recommendationServiceNaive) validateUser(ctx context.Context) (*entity.User, error) {
	userID, err := util.GetUUIDFromContext(ctx, appconstant.ContextUserID)
	if err != nil {
		return nil, err
	}
	user, err := rsn.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperror.UserNotFoundError(userID)
	}

	return user, nil
}
