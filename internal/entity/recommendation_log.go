package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type RecommendationLog struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID            uuid.UUID
	RecommendedUserID uuid.UUID
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime
}
