package services

import (
	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// PositionService represents services for positions
type PositionService struct {
	interfaces.IPositionRepository
}

// List is a service to fetch all positions
func (service *PositionService) List() (models.Positions, error) {
	positions, err := service.FindPositions()
	return positions, err
}

// FindByID is a service to fetch a position by ID
func (service *PositionService) FindByID(id int) (models.Position, error) {
	position, err := service.FindPositionByID(id)
	return position, err
}

// CreatePosition is a service to create a new position
func (service *PositionService) CreatePosition(input inputs.PositionCreateInput) (models.Position, error) {
	newPosition, err := service.StoreNewPosition(input.Output())
	return newPosition, err
}

// CreatePositions is a service to create new positions
func (service *PositionService) CreatePositions(input inputs.PositionCreateManyInput) (models.Positions, error) {
	newPositions, err := service.StoreNewPositions(input.Output())
	return newPositions, err
}

// DeleteByID is a service to delete a position by ID
func (service *PositionService) DeleteByID(id int) (models.Position, error) {
	deletedPosition, err := service.DeletePositionByID(id)
	return deletedPosition, err
}
