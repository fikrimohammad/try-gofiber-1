package interfaces

import (
	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// IClubService Docs
type IClubService interface {
	List() (models.Clubs, error)
	FindByID(id int) (models.Club, error)
	CreateClub(input inputs.ClubCreateInput) (models.Club, error)
	CreateClubs(input inputs.ClubCreateManyInput) (models.Clubs, error)
	UpdateClub(id int, input inputs.ClubUpdateInput) (models.Club, error)
	DeleteByID(id int) (models.Club, error)
}
