package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"net/http"
	"ticket-monitoring-dashboard/api"
	"ticket-monitoring-dashboard/models"
	"time"
)

func (s Server) CreateProjectProgress(c *fiber.Ctx) error {
	var req api.CreateProjectProgressJSONRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(api.Error{
			Message:   "Invalid request body",
			Error:     err.Error(),
			ErrorCode: "400",
		})
	}

	// Calculate the difference between the current time and StartTime
	duration := time.Since(*req.StartTime)

	// Calculate duration in days or hours
	var durationStr string
	if duration.Hours() >= 24 { // More than 1 day
		days := int(duration.Hours() / 24)
		durationStr = fmt.Sprintf("%d days", days)
	} else { // Less than 1 day
		hours := int(duration.Hours())
		durationStr = fmt.Sprintf("%d hours", hours)
	}

	projectResponse := models.ProjectProgress{
		Name:         *req.Name,
		ProjectID:    *req.ProjectId,
		SubStageID:   *req.SubStageId,
		CustomerName: *req.CustomerName,
		StartTime:    *req.StartTime,
		Message:      *req.Message,
		ContactEmail: *req.ContactEmail,
		Duration:     durationStr,
	}

	err := s.projectProgressService.CreateProjectProgress(context.Background(), &projectResponse)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to create project progress",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Project Progress Created",
	})
}

func (s Server) GetProjectProgressByID(c *fiber.Ctx, progressId openapi_types.UUID) error {
	projectProcess, err := s.projectProgressService.GetProjectProgressByID(context.Background(), progressId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to get project progress",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	var durationStr string

	// Check if EndTime is nil or not
	if projectProcess.EndTime.IsZero() { // Check if EndTime is zero (which is equivalent to nil for time.Time)
		// Calculate the duration using StartTime and current time (time.Now())
		duration := time.Since(projectProcess.StartTime)
		if duration.Hours() >= 24 { // More than 1 day
			days := int(duration.Hours() / 24)
			durationStr = fmt.Sprintf("%d days", days)
		} else { // Less than 1 day
			hours := int(duration.Hours())
			durationStr = fmt.Sprintf("%d hours", hours)
		}
	} else {
		// Calculate the duration using StartTime and EndTime
		duration := projectProcess.EndTime.Sub(projectProcess.StartTime)
		if duration.Hours() >= 24 { // More than 1 day
			days := int(duration.Hours() / 24)
			durationStr = fmt.Sprintf("%d days", days)
		} else { // Less than 1 day
			hours := int(duration.Hours())
			durationStr = fmt.Sprintf("%d hours", hours)
		}
	}

	projectProcess.Duration = durationStr

	return c.JSON(fiber.Map{
		"Message":         "Success",
		"ProjectProgress": projectProcess,
	})
}

func (s Server) UpdateProjectProgress(c *fiber.Ctx, progressId openapi_types.UUID) error {
	currentProgress, err := s.projectProgressService.GetProjectProgressByID(context.Background(), progressId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to get project progress",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	currentProgress.EndTime = time.Now()

	err = s.projectProgressService.UpdateProjectProgressByID(context.Background(), currentProgress)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to update project progress",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Project Progress Updated",
	})
}
