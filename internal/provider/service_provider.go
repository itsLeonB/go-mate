package provider

import (
	"github.com/itsLeonB/go-mate/internal/config"
	"github.com/itsLeonB/go-mate/internal/service"
)

type Services struct {
	Auth           service.AuthService
	Recommendation service.RecommendationService
	Subscription   service.SubscriptionService
}

func ProvideServices(
	configs *config.Config,
	repositories *Repositories,
	utils *Utils,
) *Services {
	scoringService := service.NewScoringServiceNaive()
	authService := service.NewAuthService(repositories.User, utils.Hash, utils.JWT)
	subscriptionService := service.NewSubscriptionService(repositories.Subscription, authService)

	return &Services{
		Auth: authService,
		Recommendation: service.NewRecommendationServiceNaive(
			repositories.User,
			scoringService,
			repositories.RecommendationLog,
			authService,
			subscriptionService,
		),
		Subscription: subscriptionService,
	}
}
