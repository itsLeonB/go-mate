package service

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/mapper"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/repository"
)

type recommendationServiceNaive struct {
	userRepository repository.UserRepository
	scoringService ScoringService
}

func NewRecommendationServiceNaive(userRepository repository.UserRepository, scoringService ScoringService) RecommendationService {
	return &recommendationServiceNaive{
		userRepository: userRepository,
		scoringService: scoringService,
	}
}

func (rsn *recommendationServiceNaive) GetUserRecommendations(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := rsn.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	recommendedUsers := rsn.scoringService.ScoreAndSortUsers(ctx, users)

	userResponses := make([]*model.UserResponse, len(recommendedUsers))
	for i := 0; i < len(users); i++ {
		userResponses[i] = mapper.MapUserToResponse(recommendedUsers[i])
	}

	return userResponses, nil
}
