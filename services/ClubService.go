package services

import (
	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/fikrimohammad/Premier-League-DB/models"
)

// ClubService represents services for clubs
type ClubService struct {
	interfaces.IClubRepository
}

// List is a service to fetch all clubs
func (service *ClubService) List() (models.Clubs, error) {
	clubs, err := service.FindClubs()
	return clubs, err
}

// FindByID is a service to fetch a club by ID
func (service *ClubService) FindByID(id int) (models.Club, error) {
	club, err := service.FindClubByID(id)
	return club, err
}

// CreateClub is a service to create a new club
func (service *ClubService) CreateClub(input inputs.ClubCreateInput) (models.Club, error) {
	newClub, err := service.StoreNewClub(input.Output())
	return newClub, err
}

// CreateClubs is a service to create new clubs
func (service *ClubService) CreateClubs(input inputs.ClubCreateManyInput) (models.Clubs, error) {
	newClubs, err := service.StoreNewClubs(input.Output())
	return newClubs, err
}

// UpdateClub is a service to update a club
func (service *ClubService) UpdateClub(id int, input inputs.ClubUpdateInput) (models.Club, error) {
	modifiedClub, err := service.UpdateClubByID(id, input.Output())
	return modifiedClub, err
}

// DeleteByID is a service to delete a club by ID
func (service *ClubService) DeleteByID(id int) (models.Club, error) {
	deletedClub, err := service.DeleteClubByID(id)
	return deletedClub, err
}
