package handlers

import (
	"gorm.io/gorm"
	"ticket-monitoring-dashboard/services"
)

type Server struct {
	db                     *gorm.DB
	stageServices          *services.StageService
	projectServices        *services.ProjectService
	subStageService        *services.SubStageService
	projectProgressService *services.ProjectProgressService
	customerService        *services.CustomerService
}

func NewServer(db *gorm.DB,
	stageServices *services.StageService,
	projectServices *services.ProjectService,
	subStageService *services.SubStageService,
	projectProgressService *services.ProjectProgressService,
	customerService *services.CustomerService,
) *Server {
	return &Server{
		db:                     db,
		stageServices:          stageServices,
		projectServices:        projectServices,
		subStageService:        subStageService,
		projectProgressService: projectProgressService,
		customerService:        customerService,
	}
}
