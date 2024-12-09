package provider

import (
	"github.com/itsLeonB/go-mate/internal/delivery/http/handler"
)

type Handlers struct {
	Auth           *handler.AuthHandler
	Recommendation *handler.RecommendationHandler
}

func ProvideHandlers(services *Services) *Handlers {
	return &Handlers{
		Auth: handler.NewAuthHandler(services.Auth),
	}
}
