package services

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
)

type SubStageService struct {
	subStageRepo repository.SubStageRepository
}

func NewSubStageService(subStageRepo repository.SubStageRepository) *SubStageService {
	return &SubStageService{subStageRepo: subStageRepo}
}

// Create a new substage
func (s *SubStageService) CreateSubStage(ctx context.Context, substage *models.SubStage) error {
	return s.subStageRepo.CreateSubStage(ctx, substage)
}

// Create multiple stages in batch
func (s *SubStageService) CreateSubStagesInBatch(ctx context.Context, subStages []models.SubStage) error {
	return s.subStageRepo.CreateSubStagesInBatch(ctx, subStages)
}

// Get a substage by its ID
func (s *SubStageService) GetSubStageByID(ctx context.Context, id uuid.UUID) (*models.SubStage, error) {
	return s.subStageRepo.GetSubStageByID(ctx, id)
}

// Get all subStages by a specific project ID
func (s *SubStageService) GetAllSubStagesByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.SubStage, error) {
	return s.subStageRepo.GetAllSubStageByProjectID(ctx, projectId)
}

// Get all subStages by a specific project ID
func (s *SubStageService) GetAllSubStagesByStageID(ctx context.Context, stageid uuid.UUID) ([]models.SubStage, error) {
	return s.subStageRepo.GetAllSubStageByProjectID(ctx, stageid)
}
