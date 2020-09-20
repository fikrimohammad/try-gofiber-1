package interfaces

import (
	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// IPositionService Docs
type IPositionService interface {
	List() (models.Positions, error)
	FindByID(id int) (models.Position, error)
	CreatePosition(input *inputs.PositionCreateInput) (models.Position, error)
}
