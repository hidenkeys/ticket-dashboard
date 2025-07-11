package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *models.Project) error
	GetProjectByID(ctx context.Context, id uuid.UUID) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) error
	GetAllProjects(ctx context.Context) ([]models.Project, error)
	DeleteProjectByID(ctx context.Context, id uuid.UUID) error
}
