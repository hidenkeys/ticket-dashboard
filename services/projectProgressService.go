package services

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
)

type ProjectProgressService struct {
	projectProgressRepo repository.ProjectProgressRepository
}

func NewProjectProgressService(projectProgressRepo repository.ProjectProgressRepository) *ProjectProgressService {
	return &ProjectProgressService{projectProgressRepo: projectProgressRepo}
}

// Create a new project progress record
func (s *ProjectProgressService) CreateProjectProgress(ctx context.Context, projectProgress *models.ProjectProgress) error {
	return s.projectProgressRepo.CreateProjectProgress(ctx, projectProgress)
}

// Get project progress by its ID
func (s *ProjectProgressService) GetProjectProgressByID(ctx context.Context, id uuid.UUID) (*models.ProjectProgress, error) {
	return s.projectProgressRepo.GetProjectProgressByID(ctx, id)
}

// Update a project progress record by its ID
func (s *ProjectProgressService) UpdateProjectProgressByID(ctx context.Context, projectProgress *models.ProjectProgress) error {
	return s.projectProgressRepo.UpdateProjectProgressByID(ctx, projectProgress)
}

// Get all project progress records by project ID
func (s *ProjectProgressService) GetAllProjectProgressByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.ProjectProgress, error) {
	return s.projectProgressRepo.GetAllProjectProgressByProjectID(ctx, projectId)
}

// Get all project progress records by substage ID
func (s *ProjectProgressService) GetAllProjectProgressBySubStageID(ctx context.Context, subStageId uuid.UUID) ([]models.ProjectProgress, error) {
	return s.projectProgressRepo.GetAllProjectProgressBySubStageID(ctx, subStageId)
}

// Delete project progress by ID
func (s *ProjectProgressService) DeleteProjectProgressByID(ctx context.Context, id uuid.UUID) error {
	return s.projectProgressRepo.DeleteProjectProgressByID(ctx, id)
}

func (s *ProjectProgressService) GetProjectProgressByCustomerName(ctx context.Context, customerName string) ([]models.ProjectProgress, error) {
	return s.projectProgressRepo.GetProjectProgressByCustomerName(ctx, customerName)
}
