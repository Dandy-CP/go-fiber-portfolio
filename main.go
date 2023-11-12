package main

import (
	"fmt"
	"log"

	"github.com/Dandy-CP/go-fiber-portfolio/config"
	"github.com/Dandy-CP/go-fiber-portfolio/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	setup, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	config.ConnectDB(&setup)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

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

	app.Get("*", func(c *fiber.Ctx) error {
		path := c.Path()

		if path == "/api" {
			return c.JSON(fiber.Map{
				"statusCode": 200,
				"message": "Welcome To My Portfolio Rest API",
				"version": "0.0.1",
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}