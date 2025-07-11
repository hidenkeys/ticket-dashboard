package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"ticket-monitoring-dashboard/models"
)

var DB *gorm.DB

// ConnectDatabase establishes a connection to the PostgreSQL database
func ConnectDatabase() {
	// Get the database URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("📦 Connected to database")
}

// MigrateDatabase runs migrations for the necessary models, ensuring UUID handling
func MigrateDatabase() {
	// Perform auto migration for all necessary models
	err := DB.AutoMigrate(
		&models.Stage{}, // stage model
		&models.Project{},
		&models.Customer{},
		&models.ProjectProgress{},
		&models.SubStage{},
		// Add any other models that you want to auto-migrate
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration successful")
}
