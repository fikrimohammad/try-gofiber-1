package inputs

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/gofiber/fiber/v2"
)

// ClubCreateManyInput represents parameters for creating new clubs
type ClubCreateManyInput struct {
	clubInputs []ClubCreateInput
}

// NewClubCreateManyInput is a function to initialize ClubCreateManyInput
func NewClubCreateManyInput(c *fiber.Ctx) (ClubCreateManyInput, error) {
	clubCreateManyInput := ClubCreateManyInput{}
	err := c.BodyParser(&clubCreateManyInput.clubInputs)
	return clubCreateManyInput, err
}

// Output is a function to convert ClubCreateManyInput into array of Club
func (input *ClubCreateManyInput) Output() models.Clubs {
	newClubs := models.Clubs{}
	for _, clubInput := range input.clubInputs {
		newClub := clubInput.Output()
		newClubs = append(newClubs, newClub)
	}
	return newClubs
}
