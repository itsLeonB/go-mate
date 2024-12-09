package repository

import (
	"context"

	"github.com/itsLeonB/go-mate/internal/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
}
