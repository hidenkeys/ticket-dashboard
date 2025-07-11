package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"ticket-monitoring-dashboard/api"
	"ticket-monitoring-dashboard/config"
	"ticket-monitoring-dashboard/handlers"
	"ticket-monitoring-dashboard/repository"
	"ticket-monitoring-dashboard/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}
	//var jwtSecret = []byte("your-secret-key")
	config.ConnectDatabase()
	config.MigrateDatabase()

	db := config.DB

	// Initialize repositories
	stageRepo := repository.NewStageRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	projectProgresRepo := repository.NewProjectProgressRepository(db)
	subStageRepo := repository.NewSubStageRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// Initialize services
	stageService := services.NewStageService(stageRepo)
	projectService := services.NewProjectService(projectRepo)
	projectProgressService := services.NewProjectProgressService(projectProgresRepo)
	subStageService := services.NewSubStageService(subStageRepo)
	customerService := services.NewCustomerService(customerRepo)

	server := handlers.NewServer(db, stageService, projectService, subStageService, projectProgressService, customerService)
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://crawl-app.vercel.app,https://crawl-admin.vercel.app,http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api.RegisterHandlers(app, server)

	// And we serve HTTP until the world ends.
	log.Fatal(app.Listen("0.0.0.0:8082"))
}
