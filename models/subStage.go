package models

import (
	"github.com/google/uuid"
	"time"
)

type SubStage struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StageID   uuid.UUID `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Duration  string    `gorm:"not null"` // Duration in hours or days
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
