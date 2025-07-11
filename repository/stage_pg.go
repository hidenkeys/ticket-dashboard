package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type stageRepo struct {
	db *gorm.DB
}

// NewStageRepository returns a new instance of the StageRepository implementation.
func NewStageRepository(db *gorm.DB) StageRepository {
	return &stageRepo{db: db}
}

// CreateStage inserts a new stage into the database.
func (r *stageRepo) CreateStage(ctx context.Context, stage *models.Stage) (*models.Stage, error) {
	// Create a new stage record in the database
	if err := r.db.WithContext(ctx).Create(stage).Error; err != nil {
		return nil, err
	}
	return stage, nil
}

// GetStageByID retrieves a stage by its ID from the database.
func (r *stageRepo) GetStageByID(ctx context.Context, id uuid.UUID) (*models.Stage, error) {
	var stage models.Stage
	// Query the database for a stage by its ID
	err := r.db.WithContext(ctx).First(&stage, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &stage, nil
}

// UpdateStage updates an existing stage record in the database.
func (r *stageRepo) UpdateStage(ctx context.Context, stage *models.Stage, id uuid.UUID) (*models.Stage, error) {
	// Find the stage by ID and update its details
	err := r.db.WithContext(ctx).Model(&models.Stage{}).Where("id = ?", id).Updates(stage).Error
	if err != nil {
		return nil, err
	}
	// Retrieve the updated stage
	var updatedStage models.Stage
	err = r.db.WithContext(ctx).First(&updatedStage, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &updatedStage, nil
}

// GetAllStagesByProjectID retrieves all stages associated with a given project ID.
func (r *stageRepo) GetAllStagesByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Stage, error) {
	var stages []models.Stage
	// Query the database for all stages related to the given project ID
	err := r.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&stages).Error
	if err != nil {
		return nil, err
	}
	return stages, nil
}
