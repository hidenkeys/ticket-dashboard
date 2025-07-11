package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type projectRepo struct {
	db *gorm.DB
}

// NewProjectRepository returns a new instance of ProjectRepository implementation.
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepo{db: db}
}

// CreateProject inserts a new project record into the database.
func (r *projectRepo) CreateProject(ctx context.Context, project *models.Project) error {
	// Create a new project record in the database
	if err := r.db.WithContext(ctx).Create(project).Error; err != nil {
		return err
	}
	return nil
}

// GetProjectByID retrieves a project by its ID.
func (r *projectRepo) GetProjectByID(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	var project models.Project
	// Query the database for the project by its ID
	err := r.db.WithContext(ctx).First(&project, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProject updates an existing project record in the database.
func (r *projectRepo) UpdateProject(ctx context.Context, project *models.Project) error {
	// Update the project record by its ID
	err := r.db.WithContext(ctx).Model(&models.Project{}).Where("id = ?", project.ID).Updates(project).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllProjects retrieves all projects from the database.
func (r *projectRepo) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project
	// Query the database to get all projects
	err := r.db.WithContext(ctx).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// DeleteProjectByID deletes a project by its ID.
func (r *projectRepo) DeleteProjectByID(ctx context.Context, id uuid.UUID) error {
	// Delete the project record by its ID
	err := r.db.WithContext(ctx).Delete(&models.Project{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
