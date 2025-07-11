package services

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
)

type StageService struct {
	stageRepo repository.StageRepository
}

func NewStageService(stageRepo repository.StageRepository) *StageService {
	return &StageService{stageRepo: stageRepo}
}

// Create a new stage
func (s *StageService) CreateStage(ctx context.Context, stage *models.Stage) (*models.Stage, error) {
	return s.stageRepo.CreateStage(ctx, stage)
}

// Get a stage by its ID
func (s *StageService) GetStageByID(ctx context.Context, id uuid.UUID) (*models.Stage, error) {
	return s.stageRepo.GetStageByID(ctx, id)
}

// Update an existing stage
func (s *StageService) UpdateStage(ctx context.Context, stage *models.Stage, id uuid.UUID) (*models.Stage, error) {
	return s.stageRepo.UpdateStage(ctx, stage, id)
}

// Get all stages by project ID
func (s *StageService) GetAllStagesByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Stage, error) {
	return s.stageRepo.GetAllStagesByProjectID(ctx, projectId)
}
