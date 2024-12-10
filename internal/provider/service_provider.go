package provider

import (
	"github.com/itsLeonB/go-mate/internal/config"
	"github.com/itsLeonB/go-mate/internal/service"
)

type Services struct {
	Auth           service.AuthService
	Recommendation service.RecommendationService
}

func ProvideServices(
	configs *config.Config,
	repositories *Repositories,
	utils *Utils,
) *Services {
	scoringService := service.NewScoringServiceNaive()

	return &Services{
		Auth: service.NewAuthService(repositories.User, utils.Hash, utils.JWT),
		Recommendation: service.NewRecommendationServiceNaive(
			repositories.User,
			scoringService,
			repositories.RecommendationLog,
		),
	}
}
