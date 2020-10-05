package inputs

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/gofiber/fiber/v2"
)

// PositionUpdateInput is a struct to store parameters for storing a modified position
type PositionUpdateInput struct {
	Name string `json:"name" xml:"name" form:"name"`
}

// NewPositionUpdateInput is a function to initialize PositionUpdateInput instance
func NewPositionUpdateInput(c *fiber.Ctx) (PositionUpdateInput, error) {
	positionUpdateInput := PositionUpdateInput{}
	if err := c.BodyParser(&positionUpdateInput); err != nil {
		return positionUpdateInput, err
	}
	return positionUpdateInput, nil
}

// Output is a function to convert input into a new Position model
func (input *PositionUpdateInput) Output() models.Position {
	modifiedPosition := models.Position{
		Name: input.Name,
	}
	return modifiedPosition
}
