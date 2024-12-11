package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/service"
	"github.com/itsLeonB/go-mate/tests/mocks"
	"github.com/stretchr/testify/assert"
)

type mockedInterfaces struct {
	userRepository              mocks.UserRepository
	scoringService              mocks.ScoringService
	recommendationLogRepository mocks.RecommendationLogRepository
	authService                 mocks.AuthService
	subscriptionService         mocks.SubscriptionService
}

func makeMocks(t *testing.T) *mockedInterfaces {
	return &mockedInterfaces{
		userRepository:              *mocks.NewUserRepository(t),
		scoringService:              *mocks.NewScoringService(t),
		recommendationLogRepository: *mocks.NewRecommendationLogRepository(t),
		authService:                 *mocks.NewAuthService(t),
		subscriptionService:         *mocks.NewSubscriptionService(t),
	}
}

func TestNewRecommendationService(t *testing.T) {
	t.Run("should return new recommendation service", func(t *testing.T) {
		mks := makeMocks(t)
		got := service.NewRecommendationServiceNaive(
			&mks.userRepository,
			&mks.scoringService,
			&mks.recommendationLogRepository,
			&mks.authService,
			&mks.subscriptionService,
		)

		assert.Implements(t, (*service.RecommendationService)(nil), got)
	})
}

func TestRecommendationServiceGetUserRecommendations(t *testing.T) {
	t.Run("should return recommendations", func(t *testing.T) {
		mks := makeMocks(t)
		ctx := context.Background()
		rs := service.NewRecommendationServiceNaive(
			&mks.userRepository,
			&mks.scoringService,
			&mks.recommendationLogRepository,
			&mks.authService,
			&mks.subscriptionService,
		)

		newId, _ := uuid.NewRandom()
		user := entity.User{ID: newId}

		otherId, _ := uuid.NewRandom()
		otherUser := entity.User{ID: otherId}

		userSlice := []*entity.User{&otherUser}
		log := entity.RecommendationLog{
			UserID:            user.ID,
			RecommendedUserID: otherUser.ID,
			Status:            "viewed",
		}

		mks.authService.On("ValidateUser", ctx).Return(&user, nil).Once()
		mks.recommendationLogRepository.On("FindTodayLogsByUserID", ctx, user.ID).Return([]*entity.RecommendationLog{}, nil).Once()
		mks.userRepository.On("FindAll", ctx).Return(userSlice, nil).Once()
		mks.scoringService.On("ScoreAndSortUsers", ctx, userSlice, false).Return(userSlice, nil).Once()
		mks.recommendationLogRepository.On("InsertLogs", ctx, []*entity.RecommendationLog{&log}).Return(nil).Once()

		userResponses, err := rs.GetUserRecommendations(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(userResponses))
		assert.Equal(t, otherUser.ID.String(), userResponses[0].ID)
	})
}
