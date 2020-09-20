package repositories

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"gorm.io/gorm"
)

// PositionRepository Docs
type PositionRepository struct {
	*gorm.DB
}

// FindPositions is a function to fetch all positions
func (repository *PositionRepository) FindPositions() (models.Positions, error) {
	positions := models.Positions{}
	err := repository.Find(&positions).Error
	if err != nil {
		return positions, err
	}
	return positions, nil
}

// FindPositionByID is a function to fetch a position by ID
func (repository *PositionRepository) FindPositionByID(id int) (models.Position, error) {
	position := models.Position{}
	err := repository.First(&position, id).Error
	if err != nil {
		return position, err
	}
	return position, nil
}

// StoreNewPosition is a function to store a new position
func (repository *PositionRepository) StoreNewPosition(newPosition models.Position) (models.Position, error) {
	res := repository.Create(&newPosition)
	if res.Error != nil {
		return newPosition, res.Error
	}
	return newPosition, nil
}

// StoreNewPositions is a function to store new positions
func (repository *PositionRepository) StoreNewPositions(newPositions models.Positions) (models.Positions, error) {
	res := repository.Create(&newPositions)
	if res.Error != nil {
		return newPositions, res.Error
	}
	return newPositions, nil
}
