package interfaces

import "github.com/fikrimohammad/Premier-League-DB/models"

// IPositionService Docs
type IPositionService interface {
	List() (models.Positions, error)
	ShowByID(id int) (models.Position, error)
}
