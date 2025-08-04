package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type OTPRepository interface {
	CreateOTP(ctx context.Context, otp *models.OTP) error
	MarkOTPAsUsed(ctx context.Context, id uuid.UUID) error
	GetOTPByIdentifier(ctx context.Context, email string, purpose string) (*models.OTP, error)
}
