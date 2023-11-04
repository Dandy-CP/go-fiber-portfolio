package middleware

import (
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateAuth(c *fiber.Ctx) error {
	var errors []*models.ErrorResponse
	body := new(models.User)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.Error = true
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Value() 

			errors = append(errors, &element)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": fiber.StatusBadRequest,
			"Message": "Error",
			"Data": errors,
		})
	}

	return c.Next()
}