package repositories

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"gorm.io/gorm"
)

// PositionRepository Docs
type PositionRepository struct {
	gorm.DB
}

// GetPositions is a function to fetch all positions
func (repository *PositionRepository) GetPositions() (models.Positions, error) {
	positions := models.Positions{}
	err := repository.Find(&positions).Error
	if err != nil {
		return positions, err
	}
	return positions, nil
}

// GetPositionByID is a function to fetch a position by ID
func (repository *PositionRepository) GetPositionByID(id int) (models.Position, error) {
	position := models.Position{}
	err := repository.First(&position, id).Error
	if err != nil {
		return position, err
	}
	return position, nil
}
