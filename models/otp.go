package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTP struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Identifier string    `gorm:"not null"` // could be email or phone
	Code       string    `gorm:"not null"`
	Purpose    string    `gorm:"not null"` // e.g., "email_verification", "password_reset", "pin_reset"
	ExpiresAt  time.Time `gorm:"not null"`
	IsUsed     bool      `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
