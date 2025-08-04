package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type otpRepo struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepo{db: db}
}

func (r *otpRepo) CreateOTP(ctx context.Context, otp *models.OTP) error {
	return r.db.WithContext(ctx).Create(otp).Error
}

func (r *otpRepo) MarkOTPAsUsed(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.OTP{}).Where("id = ?", id).Update("is_used", true).Error
}

func (r *otpRepo) GetOTPByIdentifier(ctx context.Context, identifier string, purpose string) (*models.OTP, error) {
	// Get the most recent OTP based on the identifier and purpose
	var otp models.OTP
	err := r.db.WithContext(ctx).Model(&models.OTP{}).
		Where("identifier = ? AND is_used = ? AND purpose = ?", identifier, false, purpose).
		Order("created_at DESC"). // Order by CreatedAt (most recent first)
		Limit(1).                 // Limit to the most recent OTP
		First(&otp).Error         // Assign the result to the 'otp' variable

	if err != nil {
		return nil, err // Return error if OTP is not found or there's a DB issue
	}

	return &otp, nil // Return the most recent OTP matching the criteria
}
