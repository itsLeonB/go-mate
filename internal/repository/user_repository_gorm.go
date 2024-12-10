package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/go-mate/internal/entity"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) UserRepository {
	return &userRepositoryGorm{db: db}
}

func (urg *userRepositoryGorm) Insert(ctx context.Context, user *entity.User) error {
	err := urg.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return eris.Wrap(err, "error inserting user")
	}

	return err
}

func (urg *userRepositoryGorm) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	err := urg.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error finding user by email")
	}

	return &user, nil
}

func (urg *userRepositoryGorm) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	err := urg.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, eris.Wrap(err, "error finding all users")
	}

	return users, nil
}

func (urg *userRepositoryGorm) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := urg.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error finding user by id")
	}

	return &user, nil
}

func (urg *userRepositoryGorm) FindByIDs(ctx context.Context, ids uuid.UUIDs) ([]*entity.User, error) {
	var users []*entity.User

	err := urg.db.WithContext(ctx).Find(&users, "id IN ?", ids).Error
	if err != nil {
		return nil, eris.Wrap(err, "error finding users by ids")
	}

	return users, nil
}
