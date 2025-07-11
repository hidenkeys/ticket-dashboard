package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type SubStageRepository interface {
	CreateSubStage(ctx context.Context, substage *models.SubStage) error
	CreateSubStagesInBatch(ctx context.Context, subStages []models.SubStage) error
	GetSubStageByID(ctx context.Context, id uuid.UUID) (*models.SubStage, error)
	GetAllSubStageByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.SubStage, error)
	GetAllSubStageByStageID(ctx context.Context, stageId uuid.UUID) ([]models.SubStage, error)
}
