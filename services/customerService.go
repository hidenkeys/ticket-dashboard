package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"regexp"
	"ticket-monitoring-dashboard/models"
	"ticket-monitoring-dashboard/repository"
)

type CustomerService struct {
	customerRepo *repository.CustomerRepo
}

func NewCustomerService(customerRepo *repository.CustomerRepo) *CustomerService {
	return &CustomerService{customerRepo: customerRepo}
}

// CreateCustomer creates a new customer after validating the input.
func (s *CustomerService) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	// Validate customer data
	if err := s.validateCustomerData(customer); err != nil {
		return err
	}

	// Check if customer already exists by email or phone number
	existingCustomer, err := s.customerRepo.GetCustomerByName(ctx, customer.Name)
	if err == nil && existingCustomer != nil {
		return errors.New("customer with this name already exists")
	}

	// Create the customer in the database
	return s.customerRepo.CreateCustomer(ctx, customer)
}

// GetCustomerByName retrieves a customer by name.
func (s *CustomerService) GetCustomerByName(ctx context.Context, name string) (*models.Customer, error) {
	// Validate the customer name format
	if len(name) == 0 {
		return nil, errors.New("customer name cannot be empty")
	}

	return s.customerRepo.GetCustomerByName(ctx, name)
}

// GetCustomersByProjectID retrieves all customers associated with a project.
func (s *CustomerService) GetCustomersByProjectID(ctx context.Context, projectId uuid.UUID) ([]models.Customer, error) {
	return s.customerRepo.GetCustomersByProjectID(ctx, projectId)
}

// validateCustomerData validates the customer's name, email, and phone number.
func (s *CustomerService) validateCustomerData(customer *models.Customer) error {
	// Validate Name
	if len(customer.Name) == 0 {
		return errors.New("customer name cannot be empty")
	}

	// Validate Email
	if err := s.validateEmail(customer.ContactEmail); err != nil {
		return err
	}

	return nil
}

// validateEmail validates the email format.
func (s *CustomerService) validateEmail(email string) error {
	// Simple email regex pattern
	const emailRegex = `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
