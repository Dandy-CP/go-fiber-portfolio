package certificatecontroller

import (
	"github.com/Dandy-CP/go-fiber-portfolio/config"
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func GetListCertificate(c *fiber.Ctx) error {
	var certificateList []models.Certificate
	limit := c.QueryInt("limit")

	if limit != 0 {
		config.DB.Limit(limit).Find(&certificateList)
	} else {
		config.DB.Find(&certificateList)
	}

	return c.Status(fiber.StatusOK).JSON(&certificateList)
}

func CreateCertificate(c *fiber.Ctx) error {
	var certificate models.Certificate

	if err := c.BodyParser(&certificate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if err := config.DB.Create(&certificate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  fiber.StatusInternalServerError,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": fiber.StatusOK,
		"Message": "Success Create Certificate",
		"Data": certificate,
	})
}

func UpdateCertificate(c *fiber.Ctx) error {
	var certificate models.Certificate
	id := c.Params("id")

	if err := c.BodyParser(&certificate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if config.DB.Where("id = ?", id).Updates(&certificate).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Success",
		"data": certificate,
	})
}

func DeleteCertificate(c *fiber.Ctx) error {
	var certificate models.Certificate
	id := c.Params("id")

	if config.DB.Where("id = ?", id).Delete(&certificate).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete",
	})
}