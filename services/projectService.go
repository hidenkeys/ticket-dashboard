package services

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
)

type ProjectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

// Create a new project
func (s *ProjectService) CreateProject(ctx context.Context, project *models.Project) error {
	return s.projectRepo.CreateProject(ctx, project)
}

// Get a project by its ID
func (s *ProjectService) GetProjectByID(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	return s.projectRepo.GetProjectByID(ctx, id)
}

// Update an existing project
func (s *ProjectService) UpdateProject(ctx context.Context, project *models.Project) error {
	return s.projectRepo.UpdateProject(ctx, project)
}

// Get all projects
func (s *ProjectService) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	return s.projectRepo.GetAllProjects(ctx)
}

// Delete a project by its ID
func (s *ProjectService) DeleteProjectByID(ctx context.Context, id uuid.UUID) error {
	return s.projectRepo.DeleteProjectByID(ctx, id)
}
