package models

import (
	"github.com/google/uuid"
	"time"
)

type ProjectProgress struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `json:"name"`
	ProjectID    uuid.UUID `gorm:"not null"`
	SubStageID   uuid.UUID `gorm:"not null"`
	CustomerName string    `json:"customer_name"`
	StartTime    time.Time `gorm:"not null"`
	Message      string    `gorm:"not null"`
	ContactEmail string    `gorm:"not null"`
	EndTime      time.Time
	Duration     string    `gorm:"not null"` // Duration in hours or days
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
