package controllers

import (
	"strconv"

	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/gofiber/fiber"
)

// PositionsController Docs
type PositionsController struct {
	interfaces.IPositionService
}

// All is a function to fetch all positions
func (controller *PositionsController) All(ctx *fiber.Ctx) {
	positions, err := controller.List()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	ctx.Status(200).JSON(fiber.Map{
		"data": positions,
	})
}

// Show is a function to fetch a position by ID
func (controller *PositionsController) Show(ctx *fiber.Ctx) {
	id, err := strconv.Atoi(ctx.Params("id"))
	position, err := controller.ShowByID(id)
	if err != nil {
		ctx.Status(404).JSON(&fiber.Map{
			"errors": err.Error(),
		})
		return
	}
	ctx.Status(200).JSON(&fiber.Map{
		"data": position,
	})
}

// func (controller *PositionsController) Create(ctx *fiber.Ctx) {
// 	// Call service to fetch all positions
// 	ctx.Status(201).JSON(fiber.Map{
// 		"data":     true,
// 		"metadata": true,
// 	})
// }

// func (controller *PositionsController) Update(ctx *fiber.Ctx) {
// 	// Call service to fetch position by params id
// 	// Update position
// 	ctx.Status(200).JSON(fiber.Map{
// 		"data":     true,
// 		"metadata": true,
// 	})
// }

// func (controller *PositionsController) Delete(ctx *fiber.Ctx) {
// 	// Call service to fetch position by params id
// 	// Delete position

// 	ctx.Status(200).JSON(fiber.Map{
// 		"data":     true,
// 		"metadata": true,
// 	})
// }
