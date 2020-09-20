package interfaces

import "github.com/fikrimohammad/Premier-League-DB/models"

// IPositionRepository Docs
type IPositionRepository interface {
	FindPositions() (models.Positions, error)
	FindPositionByID(id int) (models.Position, error)
	StoreNewPosition(position models.Position) (models.Position, error)
	StoreNewPositions(positions models.Positions) (models.Positions, error)
}
