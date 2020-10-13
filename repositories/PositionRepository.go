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
	return positions, err
}

// FindPositionByID is a function to fetch a position by ID
func (repository *PositionRepository) FindPositionByID(id int) (models.Position, error) {
	position := models.Position{}
	err := repository.First(&position, id).Error
	return position, err
}

// StoreNewPosition is a function to store a new position
func (repository *PositionRepository) StoreNewPosition(newPosition models.Position) (models.Position, error) {
	err := repository.Create(&newPosition).Error
	return newPosition, err
}

// StoreNewPositions is a function to store new positions
func (repository *PositionRepository) StoreNewPositions(newPositions models.Positions) (models.Positions, error) {
	err := repository.Create(&newPositions).Error
	return newPositions, err
}

// UpdatePositionByID is a function to update a position by ID
func (repository *PositionRepository) UpdatePositionByID(id int, modifiedPosition models.Position) (models.Position, error) {
	position, fetchErr := repository.FindPositionByID(id)
	if fetchErr != nil {
		return position, fetchErr
	}

	updateErr := repository.Model(&position).Updates(modifiedPosition).Error
	return position, updateErr
}

// DeletePositionByID is a function to delete a position by ID
func (repository *PositionRepository) DeletePositionByID(id int) (models.Position, error) {
	deletedPosition, fetchErr := repository.FindPositionByID(id)
	if fetchErr != nil {
		return deletedPosition, fetchErr
	}

	deleteErr := repository.Delete(&deletedPosition).Error
	return deletedPosition, deleteErr
}
