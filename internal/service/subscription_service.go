package service

import (
	"context"
	"time"

	"github.com/itsLeonB/go-mate/internal/appconstant"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/itsLeonB/go-mate/internal/mapper"
	"github.com/itsLeonB/go-mate/internal/model"
	"github.com/itsLeonB/go-mate/internal/repository"
)

type subscriptionServiceImpl struct {
	subscriptionRepository repository.SubscriptionRepository
	authService            AuthService
}

func NewSubscriptionService(
	subscriptionRepository repository.SubscriptionRepository,
	authService AuthService,
) SubscriptionService {
	return &subscriptionServiceImpl{
		subscriptionRepository: subscriptionRepository,
		authService:            authService,
	}
}

func (ss *subscriptionServiceImpl) AddSubscription(ctx context.Context, request *model.NewSubscriptionRequest) (*model.UserSubscriptionResponse, error) {
	user, err := ss.authService.ValidateUser(ctx)
	if err != nil {
		return nil, err
	}

	subscription := &entity.UserSubscription{
		UserID:    user.ID,
		Model:     request.Model,
		ExpiredAt: getExpiry(request.Plan),
	}

	err = ss.subscriptionRepository.Insert(ctx, subscription)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserSubscriptionToResponse(subscription), nil
}

func IsExtraRecommendation(user *entity.User) bool {
	for _, subscription := range user.Subscriptions {
		if subscription.Model == appconstant.SubscriptionExtraRecommendations && subscription.ExpiredAt.After(time.Now()) {
			return true
		}
	}

	return false
}

func getExpiry(plan string) time.Time {
	switch plan {
	case "monthly":
		return time.Now().AddDate(0, 1, 0)
	case "yearly":
		return time.Now().AddDate(1, 0, 0)
	default:
		return time.Time{}
	}
}
