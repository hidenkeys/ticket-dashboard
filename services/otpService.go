package services

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
	"time"
)

// OTPService provides logic for handling OTPs
type OTPService struct {
	OTPRepo repository.OTPRepository
}

// NewOTPService creates a new instance of OTPService
func NewOTPService(otpRepo repository.OTPRepository) *OTPService {
	return &OTPService{OTPRepo: otpRepo}
}

// CreateOTP creates a new OTP record
func (s *OTPService) CreateOTP(ctx context.Context, identifier, purpose, code string, expiresIn time.Duration) (*models.OTP, error) {
	otp := &models.OTP{
		ID:         uuid.New(),
		Identifier: identifier,
		Purpose:    purpose,
		Code:       code,
		ExpiresAt:  time.Now().UTC().Add(10 * time.Minute),
		IsUsed:     false,
	}

	err := s.OTPRepo.CreateOTP(ctx, otp)
	return otp, err
}

// MarkOTPAsUsed marks an OTP record as used
func (s *OTPService) MarkOTPAsUsed(ctx context.Context, otpID uuid.UUID) error {
	return s.OTPRepo.MarkOTPAsUsed(ctx, otpID)
}

func (s *OTPService) GetOTPByIdentifier(ctx context.Context, identifier, purpose string) (*models.OTP, error) {
	otp, err := s.OTPRepo.GetOTPByIdentifier(ctx, identifier, purpose)
	if err != nil {
		return nil, err
	}
	return otp, nil
}
