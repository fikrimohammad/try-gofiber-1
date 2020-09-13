package interfaces

import "github.com/fikrimohammad/Premier-League-DB/models"

// IPositionRepository Docs
type IPositionRepository interface {
	GetPositions() (models.Positions, error)
	GetPositionByID(id int) (models.Position, error)
}
