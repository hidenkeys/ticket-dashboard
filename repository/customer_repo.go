package repository

import (
	"context"
	"github.com/google/uuid"
	"ticket-monitoring-dashboard/models"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *models.Customer) error
	GetCustomerByName(ctx context.Context, name string) (*models.Customer, error)
	GetCustomersByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Customer, error)
}
