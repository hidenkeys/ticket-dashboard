package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type StageRepository interface {
	CreateStage(ctx context.Context, stage *models.Stage) (*models.Stage, error)
	GetStageByID(ctx context.Context, id uuid.UUID) (*models.Stage, error)
	UpdateStage(ctx context.Context, stage *models.Stage, id uuid.UUID) (*models.Stage, error)
	GetAllStagesByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Stage, error)
}
