package interfaces

import "github.com/fikrimohammad/Premier-League-DB/models"

// IClubRepository Docs
type IClubRepository interface {
	FindClubs() (models.Clubs, error)
	FindClubByID(id int) (models.Club, error)
	StoreNewClub(position models.Club) (models.Club, error)
	StoreNewClubs(positions models.Clubs) (models.Clubs, error)
	UpdateClubByID(id int, position models.Club) (models.Club, error)
	DeleteClubByID(id int) (models.Club, error)
}
