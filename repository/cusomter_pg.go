package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/models"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	err := r.db.WithContext(ctx).Create(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepo) GetCustomerByName(ctx context.Context, name string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).First(&customer, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetCustomersByProjectID fetches customers by project ID.
func (r *CustomerRepo) GetCustomersByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
