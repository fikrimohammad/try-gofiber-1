package repositories

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"gorm.io/gorm"
)

// ClubRepository Docs
type ClubRepository struct {
	*gorm.DB
}

// FindClubs is a function to fetch all clubs
func (repository *ClubRepository) FindClubs() (models.Clubs, error) {
	clubs := models.Clubs{}
	err := repository.Find(&clubs).Error
	return clubs, err
}

// FindClubByID is a function to fetch a club by ID
func (repository *ClubRepository) FindClubByID(id int) (models.Club, error) {
	club := models.Club{}
	err := repository.First(&club, id).Error
	return club, err
}

// StoreNewClub is a function to store a new club
func (repository *ClubRepository) StoreNewClub(newClub models.Club) (models.Club, error) {
	err := repository.Create(&newClub).Error
	return newClub, err
}

// StoreNewClubs is a function to store new clubs
func (repository *ClubRepository) StoreNewClubs(newClubs models.Clubs) (models.Clubs, error) {
	err := repository.Create(&newClubs).Error
	return newClubs, err
}

// UpdateClubByID is a function to update a club by ID
func (repository *ClubRepository) UpdateClubByID(id int, modifiedClub models.Club) (models.Club, error) {
	club, fetchErr := repository.FindClubByID(id)
	if fetchErr != nil {
		return club, fetchErr
	}

	updateErr := repository.Model(&club).Updates(modifiedClub).Error
	return club, updateErr
}

// DeleteClubByID is a function to delete a club by ID
func (repository *ClubRepository) DeleteClubByID(id int) (models.Club, error) {
	deletedClub, fetchErr := repository.FindClubByID(id)
	if fetchErr != nil {
		return deletedClub, fetchErr
	}

	deleteErr := repository.Delete(&deletedClub).Error
	return deletedClub, deleteErr
}
