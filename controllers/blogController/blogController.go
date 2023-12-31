package blogcontroller

import (
	"strings"

	"github.com/Dandy-CP/go-fiber-portfolio/config"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func GetListBlog(c *fiber.Ctx) error {
	var blogList []models.Blog

	valueInDB := config.DB.Find(&blogList)

	result := middleware.Pagination.With(valueInDB).Request(c.Request()).Response(&[]models.Blog{})

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetBlogDetail(c *fiber.Ctx) error {
	var blogDetail models.Blog
	id := c.Params("id")

	if config.DB.Where("id = ?", id).Find(&blogDetail).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&blogDetail)
}

func GetBlogByTitle(c *fiber.Ctx) error {
	var blogDetail models.Blog
	blogTitle := c.Params("title")

	valueTitle := strings.Replace(blogTitle, "-", " ", -1)

	if config.DB.Where("title = ?", valueTitle).Find(&blogDetail).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&blogDetail)
}

func CreateBlog(c *fiber.Ctx) error {
	var blog models.Blog

	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if err := config.DB.Create(&blog).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  fiber.StatusInternalServerError,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": fiber.StatusOK,
		"Message": "Success Create Blog",
		"Data": blog,
	})
}

func UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog

	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if config.DB.Where("id = ?", id).Updates(&blog).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Success",
		"data": blog,
	})
}

func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog

	if config.DB.Where("id = ?", id).Delete(&blog).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  fiber.StatusNotFound,
			"message": "Data Not Found",
		})
	}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete",
	})
}