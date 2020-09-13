package configs

import (
	"sync"

	"github.com/fikrimohammad/Premier-League-DB/controllers"
	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/fikrimohammad/Premier-League-DB/repositories"
	"github.com/fikrimohammad/Premier-League-DB/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IServiceContainer interface {
	InjectPositionsController() controllers.PositionsController
}

type kernel struct{}

func (k *kernel) InjectPositionsController() controllers.PositionsController {
	dsn := "user=postgres password=postgres dbname=pl_development port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Position{})

	positionRepository := &repositories.PositionRepository{*db}
	positionService := &services.PositionService{positionRepository}
	positionsController := controllers.PositionsController{positionService}

	return positionsController
}

var (
	k             *kernel
	containerOnce sync.Once
)

// ServiceContainer is a function to inject dependency
func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
