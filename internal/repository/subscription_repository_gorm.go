package repository

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type subscriptionRepositoryGorm struct {
	db *gorm.DB
}

func NewSubscriptionRepositoryGorm(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepositoryGorm{db}
}

func (srg *subscriptionRepositoryGorm) Insert(ctx context.Context, subscription *entity.UserSubscription) error {
	err := srg.db.Create(subscription).Error
	if err != nil {
		return eris.Wrap(err, "error while inserting subscription")
	}

	return nil
}
