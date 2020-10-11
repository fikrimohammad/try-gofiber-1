package inputs

import (
	"time"

	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/gofiber/fiber/v2"
)

// ClubCreateInput is a struct to store parameters for storing a new club
type ClubCreateInput struct {
	FullName     string             `json:"full_name" xml:"full_name" form:"full_name"`
	ShortName    string             `json:"short_name" xml:"short_name" form:"short_name"`
	Abbreviation string             `json:"abbreviation" xml:"abbreviation" form:"abbreviation"`
	IconURL      string             `json:"icon_url" xml:"icon_url" form:"icon_url"`
	WebURL       string             `json:"web_url" xml:"web_url" form:"web_url"`
	Stadium      StadiumCreateInput `json:"stadium" xml:"stadium" form:"stadium"`
}

// StadiumCreateInput is a struct to store parameters for storing the club's stadium
type StadiumCreateInput struct {
	Name        string    `json:"name" xml:"name" form:"name"`
	Description string    `json:"description" xml:"description" form:"description"`
	Capacity    int       `json:"capacity" xml:"capacity" form:"capacity"`
	BuiltOn     time.Time `json:"built_on" xml:"built_on" form:"built_on"`
}

// NewClubCreateInput is a function to initialize ClubCreateInput instance
func NewClubCreateInput(c *fiber.Ctx) (ClubCreateInput, error) {
	clubCreateInput := ClubCreateInput{}
	if err := c.BodyParser(&clubCreateInput); err != nil {
		return clubCreateInput, err
	}
	return clubCreateInput, nil
}

// Output is a function to convert input into a new Club model
func (input *ClubCreateInput) Output() models.Club {
	newClub := models.Club{
		FullName:     input.FullName,
		ShortName:    input.ShortName,
		Abbreviation: input.Abbreviation,
		IconURL:      input.IconURL,
		WebURL:       input.WebURL,
		Stadium: models.Stadium{
			Name:        input.Stadium.Name,
			Description: input.Stadium.Description,
			BuiltOn:     input.Stadium.BuiltOn,
			Capacity:    input.Stadium.Capacity,
		},
	}
	return newClub
}
