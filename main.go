package main

import (
	"github.com/Dandy-CP/go-fiber-portfolio/controllers/myprojectscontroller"
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	myProjects := api.Group("/my-projects")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"statusCode": 200,
			"message": "Welcome To My Portfolio Rest API",
			"version": "0.0.1",
		})
	})

	myProjects.Get("/", myprojectscontroller.GetProjects)

	myProjects.Post("/", 
		models.ValidateMyProjects, 
			myprojectscontroller.CreateProjects)

	myProjects.Put("/:id",
		models.ValidateMyProjects,
			myprojectscontroller.UpdateProjects)

	myProjects.Delete("/:id", myprojectscontroller.DeleteProjects)

	app.Listen(":8000")
}