package inputs

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/gofiber/fiber/v2"
)

// PositionCreateInput is a struct to store parameters for storing a new position
type PositionCreateInput struct {
	Name string `json:"name" xml:"name" form:"name"`
}

// NewPositionCreateInput is a function to initialize PositionCreateInput instance
func NewPositionCreateInput(c *fiber.Ctx) (PositionCreateInput, error) {
	positionCreateInput := PositionCreateInput{}
	err := c.BodyParser(&positionCreateInput)
	return positionCreateInput, err
}

// Output is a function to convert input into a new Position model
func (input *PositionCreateInput) Output() models.Position {
	newPosition := models.Position{
		Name: input.Name,
	}
	return newPosition
}
