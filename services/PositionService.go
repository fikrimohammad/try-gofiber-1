package services

import (
	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// PositionService Docs
type PositionService struct {
	interfaces.IPositionRepository
}

// List is a function to fetch all positions
func (service *PositionService) List() (models.Positions, error) {
	positions, err := service.FindPositions()
	return positions, err
}

// FindByID is a function to fetch a position by ID
func (service *PositionService) FindByID(id int) (models.Position, error) {
	position, err := service.FindPositionByID(id)
	return position, err
}

// Create is function to create a new position
func (service *PositionService) CreatePosition(input *inputs.PositionCreateInput) (models.Position, error) {
	position, err := service.StoreNewPosition(input.Output())
	return position, err
}
