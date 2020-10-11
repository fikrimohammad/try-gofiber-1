package configs

import (
	"sync"

	"github.com/fikrimohammad/Premier-League-DB/controllers"
	"github.com/fikrimohammad/Premier-League-DB/database"
	"github.com/fikrimohammad/Premier-League-DB/repositories"
	"github.com/fikrimohammad/Premier-League-DB/services"
)

// IServiceContainer represents ServiceContainer instance
type IServiceContainer interface {
	InjectPositionsController() controllers.PositionsController
	InjectClubsController() controllers.ClubsController
}

type kernel struct{}

func (k *kernel) InjectPositionsController() controllers.PositionsController {
	positionRepository := &repositories.PositionRepository{database.Instance()}
	positionService := &services.PositionService{positionRepository}
	positionsController := controllers.PositionsController{positionService}

	return positionsController
}

func (k *kernel) InjectClubsController() controllers.ClubsController {
	clubRepository := &repositories.ClubRepository{database.Instance()}
	clubService := &services.ClubService{clubRepository}
	clubsController := controllers.ClubsController{clubService}

	return clubsController
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
