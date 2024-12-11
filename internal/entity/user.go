package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	Subscriptions []*UserSubscription `gorm:"foreignKey:UserID"`
}

type UserSubscription struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID
	Model     string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
	DeletedAt sql.NullTime
}
