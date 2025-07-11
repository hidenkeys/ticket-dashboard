package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type subStageRepo struct {
	db *gorm.DB
}

// NewSubStageRepository returns a new instance of SubStageRepository implementation.
func NewSubStageRepository(db *gorm.DB) SubStageRepository {
	return &subStageRepo{db: db}
}

// CreateSubStage inserts a new substage record into the database.
func (r *subStageRepo) CreateSubStage(ctx context.Context, substage *models.SubStage) error {
	// Create a new substage record in the database
	if err := r.db.WithContext(ctx).Create(substage).Error; err != nil {
		return err
	}
	return nil
}

// CreateSubStagesInBatch inserts multiple substage records in one transaction.
func (r *subStageRepo) CreateSubStagesInBatch(ctx context.Context, subStages []models.SubStage) error {
	// Insert all substage records in batch within a single transaction
	if err := r.db.WithContext(ctx).Create(&subStages).Error; err != nil {
		return err
	}
	return nil
}

// GetSubStageByID retrieves a substage record by its ID.
func (r *subStageRepo) GetSubStageByID(ctx context.Context, id uuid.UUID) (*models.SubStage, error) {
	var substage models.SubStage
	// Query the database for a substage by its ID
	err := r.db.WithContext(ctx).First(&substage, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &substage, nil
}

// GetAllSubStageByProjectID retrieves all substage records related to a given project ID.
func (r *subStageRepo) GetAllSubStageByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.SubStage, error) {
	var subStages []models.SubStage
	// Query the database for all substage records related to the project ID
	err := r.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&subStages).Error
	if err != nil {
		return nil, err
	}
	return subStages, nil
}

func (r *subStageRepo) GetAllSubStageByStageID(ctx context.Context, stageid uuid.UUID) ([]models.SubStage, error) {
	var subStages []models.SubStage
	// Query the database for all substage records related to the project ID
	err := r.db.WithContext(ctx).Where("stage_id = ?", stageid).Find(&subStages).Error
	if err != nil {
		return nil, err
	}
	return subStages, nil
}
