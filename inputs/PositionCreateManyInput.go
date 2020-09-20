package inputs

import (
	"github.com/fikrimohammad/Premier-League-DB/models"
	"github.com/gofiber/fiber/v2"
)

// PositionCreateManyInput represents parameters for creating new positions
type PositionCreateManyInput struct {
	Inputs []PositionCreateInput
}

// NewPositionCreateManyInput is a function to initialize PositionCreateManyInput
func NewPositionCreateManyInput(c *fiber.Ctx) (PositionCreateManyInput, error) {
	positionCreateManyInput := PositionCreateManyInput{}
	if err := c.BodyParser(&positionCreateManyInput.Inputs); err != nil {
		return positionCreateManyInput, err
	}
	return positionCreateManyInput, nil
}

// Output is a function to convert PositionCreateManyInput into array of Position
func (input *PositionCreateManyInput) Output() models.Positions {
	newPositions := models.Positions{}
	for _, inputValue := range input.Inputs {
		newPosition := inputValue.Output()
		newPositions = append(newPositions, newPosition)
	}
	return newPositions
}
