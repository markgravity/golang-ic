package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"unique;primaryKey;type:uuid;default:uuid_generate_v1()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
