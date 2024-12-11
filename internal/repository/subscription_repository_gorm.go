package repository

import (
	"context"

	"github.com/google/uuid"
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

func (srg *subscriptionRepositoryGorm) FindByUserIDandModel(
	ctx context.Context,
	userID uuid.UUID,
	model string,
) ([]*entity.UserSubscription, error) {
	var subscriptions []*entity.UserSubscription

	err := srg.db.Where("user_id = ? AND model = ?", userID, model).Find(&subscriptions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error while finding subscription by user id and model")
	}

	return subscriptions, nil
}
