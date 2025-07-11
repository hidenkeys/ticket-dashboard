package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type projectProgressRepo struct {
	db *gorm.DB
}

// NewProjectProgressRepository returns a new instance of ProjectProgressRepository implementation.
func NewProjectProgressRepository(db *gorm.DB) ProjectProgressRepository {
	return &projectProgressRepo{db: db}
}

// CreateProjectProgress inserts a new project progress record into the database.
func (r *projectProgressRepo) CreateProjectProgress(ctx context.Context, project *models.ProjectProgress) error {
	// Create a new project progress record in the database
	if err := r.db.WithContext(ctx).Create(project).Error; err != nil {
		return err
	}
	return nil
}

// GetProjectProgressByID retrieves the project progress by its ID.
func (r *projectProgressRepo) GetProjectProgressByID(ctx context.Context, id uuid.UUID) (*models.ProjectProgress, error) {
	var projectProgress models.ProjectProgress
	// Query the database for the project progress by its ID
	err := r.db.WithContext(ctx).First(&projectProgress, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &projectProgress, nil
}

// UpdateProjectProgressByID updates an existing project progress record by its ID.
func (r *projectProgressRepo) UpdateProjectProgressByID(ctx context.Context, project *models.ProjectProgress) error {
	// Find the project progress by its ID and update the record
	err := r.db.WithContext(ctx).Model(&models.ProjectProgress{}).Where("id = ?", project.ID).Updates(project).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllProjectProgressByProjectID retrieves all project progress records related to a given project ID.
func (r *projectProgressRepo) GetAllProjectProgressByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.ProjectProgress, error) {
	var projects []models.ProjectProgress
	// Query the database for all project progress records related to the given project ID
	err := r.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// GetAllProjectProgressBySubStageID retrieves all project progress records related to a given substage ID.
func (r *projectProgressRepo) GetAllProjectProgressBySubStageID(ctx context.Context, subStageId uuid.UUID) ([]models.ProjectProgress, error) {
	var projects []models.ProjectProgress
	// Query the database for all project progress records related to the given substage ID
	err := r.db.WithContext(ctx).Where("sub_stage_id = ?", subStageId).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// DeleteProjectProgressByID deletes a project progress record by its ID.
func (r *projectProgressRepo) DeleteProjectProgressByID(ctx context.Context, id uuid.UUID) error {
	// Delete the project progress record from the database
	err := r.db.WithContext(ctx).Delete(&models.ProjectProgress{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *projectProgressRepo) GetProjectProgressByCustomerName(ctx context.Context, customerName string) ([]models.ProjectProgress, error) {
	var projects []models.ProjectProgress
	// Query the database for all project progress records related to the given customer name
	err := r.db.WithContext(ctx).Where("customer_name = ?", customerName).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
