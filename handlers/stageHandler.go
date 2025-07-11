package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"log"
	"net/http"
	"ticket-monitoring-dashboard/api"
	"ticket-monitoring-dashboard/models"
	"time"
)

func (s Server) CreateStage(c *fiber.Ctx) error {
	var req api.CreateStageJSONRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(api.Error{
			Message:   "invalid request body",
			Error:     err.Error(),
			ErrorCode: "400",
		})
	}

	createdStage := models.Stage{
		Name:        *req.Name,
		Description: *req.Description,
		ProjectID:   *req.ProjectId,
	}

	stage, err := s.stageServices.CreateStage(context.Background(), &createdStage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to create stage",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Created stage",
		"stage":   stage,
	})
}

func (s Server) GetStageByID(c *fiber.Ctx, stageId openapi_types.UUID) error {
	stage, err := s.stageServices.GetStageByID(context.Background(), stageId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to get stage",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success getting stage",
		"stage":   stage,
	})
}

func (s Server) UpdateStage(c *fiber.Ctx, stageId openapi_types.UUID) error {
	var req api.UpdateStageJSONRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(api.Error{
			Message:   "invalid request body",
			Error:     err.Error(),
			ErrorCode: "400",
		})
	}

	createdStage := models.Stage{
		Name:        *req.Name,
		Description: *req.Description,
	}

	stage, err := s.stageServices.UpdateStage(context.Background(), &createdStage, stageId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Failed to create stage",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success updating stage",
		"stage":   stage,
	})
}

func (s Server) BatchCreateSubStages(c *fiber.Ctx) error {
	// Create the request struct to hold the parsed body
	var req api.BatchCreateSubStagesJSONRequestBody

	// Parse the request body into the req struct
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	// Check if SubStages are provided in the request
	if len(req) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No sub-stages provided",
		})
	}

	// Convert the req (api.BatchCreateSubStagesJSONRequestBody) to []models.SubStage
	var subStages []models.SubStage
	for _, substageAPI := range req {
		// Calculate the difference between the current time and StartTime
		duration := time.Since(*substageAPI.StartTime)

		// Calculate duration in days or hours
		var durationStr string
		if duration.Hours() >= 24 { // More than 1 day
			days := int(duration.Hours() / 24)
			durationStr = fmt.Sprintf("%d days", days)
		} else { // Less than 1 day
			hours := int(duration.Hours())
			durationStr = fmt.Sprintf("%d hours", hours)
		}
		// Convert api.SubStage to models.SubStage
		substage := models.SubStage{
			Name:      *substageAPI.Name,
			Duration:  durationStr,
			StageID:   *substageAPI.StageId, // Assuming StageId from API is the same as StageID in the models
			StartTime: time.Now(),           // Set current time
		}

		// Append the converted substage to the subStages slice
		subStages = append(subStages, substage)
	}

	// Call the CreateSubStagesInBatch method from SubStageService
	err := s.subStageService.CreateSubStagesInBatch(c.Context(), subStages)
	if err != nil {
		log.Printf("Error creating sub-stages: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create sub-stages",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sub-stages created successfully",
		"data":    subStages,
	})
}
