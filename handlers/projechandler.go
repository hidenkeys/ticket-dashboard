package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"net/http"
	"strings"
	"ticket-monitoring-dashboard/api"
	"ticket-monitoring-dashboard/models"
)

func ProcessString(input string) string {
	// Step 1: Trim the leading and trailing spaces
	trimmed := strings.TrimSpace(input)

	// Step 2: Replace any space between words with an underscore
	result := strings.ReplaceAll(trimmed, " ", "_")

	return result
}

func (s Server) CreateCustomer(c *fiber.Ctx) error {
	var req api.CreateCustomerJSONRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(api.Error{
			Message:   "invalid request body",
			Error:     err.Error(),
			ErrorCode: "400",
		})
	}

	if req.Name != nil {
		*req.Name = ProcessString(*req.Name)
	}
	customer := models.Customer{
		ProjectId:     *req.ProjectId,
		Name:          *req.Name,
		ContactEmail:  *req.ContactEmail,
		ContactPerson: *req.ContactPerson,
	}
	err := s.customerService.CreateCustomer(context.Background(), &customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot create customer",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "customer created",
	})

}

func (s Server) GetAllProjects(c *fiber.Ctx) error {
	projects, err := s.projectServices.GetAllProjects(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot get all projects",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message":  "Success",
		"Projects": projects,
	})

}

func (s Server) CreateProject(c *fiber.Ctx) error {
	var req api.CreateProjectJSONRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(api.Error{
			Message:   "invalid request body",
			Error:     err.Error(),
			ErrorCode: "400",
		})
	}
	project := models.Project{
		Name:     *req.Name,
		Customer: *req.Customer,
	}

	err := s.projectServices.CreateProject(context.Background(), &project)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot create project",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "project created",
	})
}

func (s Server) GetProjectByID(c *fiber.Ctx, projectId openapi_types.UUID) error {
	// Fetch project data
	project, err := s.projectServices.GetProjectByID(context.Background(), projectId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot get project",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	// Fetch stages related to the project
	stages, err := s.stageServices.GetAllStagesByProjectID(context.Background(), projectId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot get stages",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	// Fetch substages for each stage
	var stagesWithSubStages []map[string]interface{}
	for _, stage := range stages {
		// Fetch substages related to each stage
		substages, err := s.subStageService.GetAllSubStagesByStageID(context.Background(), stage.ID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(api.Error{
				Message:   "Cannot get substages",
				Error:     err.Error(),
				ErrorCode: "500",
			})
		}

		// Add stage and its substages to the response structure
		stagesWithSubStages = append(stagesWithSubStages, map[string]interface{}{
			"stage":     stage,     // Add stage data
			"substages": substages, // Add substages for this stage
		})
	}

	// Fetch customers related to the project
	customers, err := s.customerService.GetCustomersByProjectID(context.Background(), projectId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(api.Error{
			Message:   "Cannot get customers",
			Error:     err.Error(),
			ErrorCode: "500",
		})
	}

	// For each customer, fetch their project progress
	var customersWithProjectProgress []map[string]interface{}
	for _, customer := range customers {
		projectProgress, err := s.projectProgressService.GetProjectProgressByCustomerName(context.Background(), customer.Name)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(api.Error{
				Message:   "Cannot get project progress",
				Error:     err.Error(),
				ErrorCode: "500",
			})
		}

		// Add customer and their project progress to the response structure
		customersWithProjectProgress = append(customersWithProjectProgress, map[string]interface{}{
			"customer":        customer,        // Customer data
			"projectProgress": projectProgress, // Project progress data for this customer
		})
	}

	// Create the final structured response
	response := map[string]interface{}{
		"project":   project,                      // Add the main project data
		"stages":    stagesWithSubStages,          // Add stages and their substages
		"customers": customersWithProjectProgress, // Add customers with their project progress
	}

	// Return the structured response
	return c.Status(http.StatusOK).JSON(response)
}
