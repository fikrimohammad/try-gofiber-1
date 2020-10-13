package controllers

import (
	"strconv"

	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/gofiber/fiber/v2"
)

// PositionsController Docs
type PositionsController struct {
	interfaces.IPositionService
}

// All is an API to show all positions
func (controller *PositionsController) All(c *fiber.Ctx) error {
	positions, err := controller.List()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": positions,
	})
}

// Show is an API to show position by ID
func (controller *PositionsController) Show(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	position, err := controller.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": position,
	})
}

// Create is an API to create a position
func (controller *PositionsController) Create(c *fiber.Ctx) error {
	positionCreateinput, inputErr := inputs.NewPositionCreateInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	newPosition, err := controller.CreatePosition(positionCreateinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": newPosition,
	})
}

// CreateMany is an API to create positions
func (controller *PositionsController) CreateMany(c *fiber.Ctx) error {
	positionCreateManyinput, inputErr := inputs.NewPositionCreateManyInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	newPositions, err := controller.CreatePositions(positionCreateManyinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": newPositions,
	})
}

// Update is an API to update a position
func (controller *PositionsController) Update(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	positionUpdateInput, inputErr := inputs.NewPositionUpdateInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	modifiedPosition, err := controller.UpdatePosition(id, positionUpdateInput)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": modifiedPosition,
	})
}

// Delete is an API to delete a position by ID
func (controller *PositionsController) Delete(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	deletedPosition, err := controller.DeleteByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": deletedPosition,
	})
}
