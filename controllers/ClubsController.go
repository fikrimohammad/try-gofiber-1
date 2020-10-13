package controllers

import (
	"strconv"

	"github.com/fikrimohammad/Premier-League-DB/inputs"
	"github.com/fikrimohammad/Premier-League-DB/interfaces"
	"github.com/gofiber/fiber/v2"
)

// ClubsController Docs
type ClubsController struct {
	interfaces.IClubService
}

// All is an API to show all clubs
func (controller *ClubsController) All(c *fiber.Ctx) error {
	clubs, err := controller.List()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": clubs,
	})
}

// Show is an API to show club by ID
func (controller *ClubsController) Show(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	club, err := controller.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": club,
	})
}

// Create is an API to create a club
func (controller *ClubsController) Create(c *fiber.Ctx) error {
	clubCreateinput, inputErr := inputs.NewClubCreateInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	newClub, err := controller.CreateClub(clubCreateinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": newClub,
	})
}

// CreateMany is an API to create clubs
func (controller *ClubsController) CreateMany(c *fiber.Ctx) error {
	clubCreateManyinput, inputErr := inputs.NewClubCreateManyInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	newClubs, err := controller.CreateClubs(clubCreateManyinput)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": newClubs,
	})
}

// Update is an API to update a club
func (controller *ClubsController) Update(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	clubUpdateInput, inputErr := inputs.NewClubUpdateInput(c)
	if inputErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": inputErr.Error(),
		})
	}

	modifiedClub, err := controller.UpdateClub(id, clubUpdateInput)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": modifiedClub,
	})
}

// Delete is an API to delete a club by ID
func (controller *ClubsController) Delete(c *fiber.Ctx) error {
	id, parseIntErr := strconv.Atoi(c.Params("id"))
	if parseIntErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": parseIntErr.Error(),
		})
	}

	deletedClub, err := controller.DeleteByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": deletedClub,
	})
}
