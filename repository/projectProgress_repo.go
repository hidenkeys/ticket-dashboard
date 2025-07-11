package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type ProjectProgressRepository interface {
	GetProjectProgressByID(ctx context.Context, id uuid.UUID) (*models.ProjectProgress, error)
	UpdateProjectProgressByID(ctx context.Context, project *models.ProjectProgress) error
	GetAllProjectProgressByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.ProjectProgress, error)
	GetAllProjectProgressBySubStageID(ctx context.Context, subStageId uuid.UUID) ([]models.ProjectProgress, error)
	CreateProjectProgress(ctx context.Context, project *models.ProjectProgress) error
	DeleteProjectProgressByID(ctx context.Context, id uuid.UUID) error
	GetProjectProgressByCustomerName(ctx context.Context, customerName string) ([]models.ProjectProgress, error)
}
