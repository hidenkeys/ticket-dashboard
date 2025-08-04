package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"ticket-monitoring-dashboard/api"
	"ticket-monitoring-dashboard/utils"
	"time"
)

func (s Server) PostAuthRequestOtp(c *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}
	var req api.PostAuthRequestOtpJSONBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Generate OTP
	otpCode := utils.GenerateOTP() // Assume a utility function to generate OTP
	_, err := s.otpService.CreateOTP(context.Background(), req.Identifier, "telegram-verification", otpCode, time.Minute*10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create OTP"})
	}

	// Send OTP via SMS/Email (implement the sending logic in your utility)
	// 🚀 Call Python script to send OTP via Telegram
	apiID := os.Getenv("APIID")
	apiHash := os.Getenv("APIHASH")
	phoneNumber := req.Identifier // Make sure it’s in international format
	message := fmt.Sprintf("Your OTP is: %s", otpCode)

	cmd := exec.Command("python3", "test4.py", apiID, apiHash, phoneNumber, message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to send OTP via Telegram",
			"details": string(output),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP sent"})
}

func (s Server) PostAuthVerifyOtp(c *fiber.Ctx) error {
	var req api.PostAuthVerifyOtpJSONBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Retrieve the OTP for validation
	otp, err := s.otpService.GetOTPByIdentifier(context.Background(), req.Identifier, "telegram-verification")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "OTP not found"})
	}

	// Check if the OTP has expired
	if otp.ExpiresAt.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "OTP has expired"})
	}

	// Compare the OTP provided with the stored OTP
	if req.Otp != otp.Code {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	// Mark OTP as used
	err = s.otpService.MarkOTPAsUsed(context.Background(), otp.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to mark OTP as used"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP validated successfully"})
}
