package provider

import (
	"github.com/itsLeonB/go-mate/internal/delivery/http/handler"
)

type Handlers struct {
	Auth           *handler.AuthHandler
	Recommendation *handler.RecommendationHandler
	Subscription   *handler.SubscriptionHandler
}

func ProvideHandlers(services *Services) *Handlers {
	return &Handlers{
		Auth:           handler.NewAuthHandler(services.Auth),
		Recommendation: handler.NewRecommendationHandler(services.Recommendation),
		Subscription:   handler.NewSubscriptionHandler(services.Subscription),
	}
}
