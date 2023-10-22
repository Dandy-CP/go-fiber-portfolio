package myprojectscontroller

import (
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	var myProjects []models.MyProjects

	models.DB.Find(&myProjects)

	return c.Status(fiber.StatusOK).JSON(myProjects)
}

func CreateProjects(c *fiber.Ctx) error {
	var myProjects models.MyProjects

	if err := c.BodyParser(&myProjects); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&myProjects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(myProjects)
}

func UpdateProjects(c *fiber.Ctx) error {
	return nil
}

func DeleteProjects(c *fiber.Ctx) error {
	return nil
}