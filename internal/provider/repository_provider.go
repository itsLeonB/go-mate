package provider

import (
	"github.com/itsLeonB/go-mate/internal/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User              repository.UserRepository
	RecommendationLog repository.RecommendationLogRepository
	Subscription      repository.SubscriptionRepository
}

func ProvideRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:              repository.NewUserRepositoryGorm(db),
		RecommendationLog: repository.NewRecommendationLogRepositoryGorm(db),
		Subscription:      repository.NewSubscriptionRepositoryGorm(db),
	}
}
