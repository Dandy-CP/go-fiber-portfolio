package myprojectscontroller

import (
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	var myProjects []models.MyProjects

	limit := c.QueryInt("limit")

	if limit != 0 {
		models.DB.Limit(limit).Find(&myProjects)
	} else {
		models.DB.Find(&myProjects)
	}

	return c.Status(fiber.StatusOK).JSON(&myProjects)
}

func CreateProjects(c *fiber.Ctx) error {
	var myProjects models.MyProjects

	if err := models.DB.Create(&myProjects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  fiber.StatusInternalServerError,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	return c.JSON(myProjects)
}

func UpdateProjects(c *fiber.Ctx) error {
	id := c.Params("id")
	var myProjects models.MyProjects

	if models.DB.Where("id = ?", id).Updates(&myProjects).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Success",
		"data": myProjects,
	})
}

func DeleteProjects(c *fiber.Ctx) error {
	id := c.Params("id")
	var myProjects models.MyProjects

	if models.DB.Where("id = ?", id).Delete(&myProjects).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete",
	})
}