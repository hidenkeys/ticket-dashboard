package models

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectId     uuid.UUID `gorm:"not null"`
	Name          string    `gorm:"not null"`
	ContactPerson string    `gorm:"not null"`
	ContactEmail  string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp"`
}
