package main

import (
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/Dandy-CP/go-fiber-portfolio/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	myProjects := api.Group("/my-projects")
	certificate := api.Group("/certificate")
	blog := api.Group("/blog")
	auth := api.Group("/auth")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"statusCode": 200,
			"message": "Welcome To My Portfolio Rest API",
			"version": "0.0.1",
		})
	})

	routes.MyProjectsRoutes(myProjects)
	routes.BlogRoutes(blog)
	routes.CertificateRoutes(certificate)
	routes.AuthRoutes(auth)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}