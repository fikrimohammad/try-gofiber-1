package services

import (
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// PositionService Docs
type PositionService struct {
	interfaces.IPositionRepository
}

// List is a function to fetch all positions
func (service *PositionService) List() (models.Positions, error) {
	positions, err := service.GetPositions()
	return positions, err
}

// ShowByID is a function to fetch a position by ID
func (service *PositionService) ShowByID(id int) (models.Position, error) {
	position, err := service.GetPositionByID(id)
	return position, err
}
