package main

import (
	blogcontroller "github.com/Dandy-CP/go-fiber-portfolio/controllers/blogController"
	"github.com/Dandy-CP/go-fiber-portfolio/controllers/myprojectscontroller"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	myProjects := api.Group("/my-projects")
	blog := api.Group("/blog")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"statusCode": 200,
			"message": "Welcome To My Portfolio Rest API",
			"version": "0.0.1",
		})
	})

	myProjects.Get("/", myprojectscontroller.GetProjects)
	myProjects.Post("/", middleware.ValidateMyProjects, myprojectscontroller.CreateProjects)
	myProjects.Put("/:id", middleware.ValidateMyProjects, myprojectscontroller.UpdateProjects)
	myProjects.Delete("/:id", myprojectscontroller.DeleteProjects)

	blog.Get("/", blogcontroller.GetListBlog)
	blog.Get("/:id", blogcontroller.GetBlogDetail)
	blog.Post("/", middleware.ValidateBlog, blogcontroller.CreateBlog)
	blog.Put("/:id", middleware.ValidateBlog, blogcontroller.UpdateBlog)
	blog.Delete("/:id", blogcontroller.DeleteBlog)


	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}