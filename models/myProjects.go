package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	Base struct {
		ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"update_at"`
		DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	}

	MyProjects struct {
		Base
		ProjectName 		string					`validate:"required,min=3"`
		Description 		string					`validate:"required,min=3"`
		TechStack 			pq.StringArray	`gorm:"type:text[];" validate:"required,min=1,dive"`
		DemoLink 				string					`validate:"required"`
		SourceCodeLink 	string					`validate:"required"`
		ProjectImage 		string					`validate:"required"`
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
}
)

func (value *Base) BeforeCreate(tx *gorm.DB) (err error) {
	value.ID = uuid.NewV4()
	return
}

var Validator = validator.New()
func ValidateMyProjects(c *fiber.Ctx) error {
    var errors []*ErrorResponse
    body := new(MyProjects)

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
				var element ErrorResponse
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