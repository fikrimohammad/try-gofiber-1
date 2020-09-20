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
			"errors": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": positions,
	})
}

// Show is an API to show position by ID
func (controller *PositionsController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	position, err := controller.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
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
			"errors": inputErr,
		})
	}

	newPosition, err := controller.CreatePosition(positionCreateinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"errors": err,
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
			"errors": inputErr,
		})
	}

	newPositions, err := controller.CreatePositions(positionCreateManyinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"errors": err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": newPositions,
	})
}

// func (controller *PositionsController) Update(c *fiber.Ctx) {
// 	// Call service to fetch position by params id
// 	// Update position
// 	c.Status(200).JSON(fiber.Map{
// 		"data":     true,
// 		"metadata": true,
// 	})
// }

// Delete is an API to delete a position by ID
func (controller *PositionsController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	deletedPosition, err := controller.DeleteByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": deletedPosition,
	})
}
