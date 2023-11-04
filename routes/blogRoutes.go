package routes

import (
	blogcontroller "github.com/Dandy-CP/go-fiber-portfolio/controllers/blogController"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(blog fiber.Router) {
	blog.Get("/", blogcontroller.GetListBlog)

	blog.Get("/:id", blogcontroller.GetBlogDetail)

	blog.Post("/",
	middleware.AuthGuard,
		middleware.ValidateBlog,
			blogcontroller.CreateBlog)

	blog.Put("/:id",
	middleware.AuthGuard,
		middleware.ValidateBlog,
			blogcontroller.UpdateBlog)

	blog.Delete("/:id",
		middleware.AuthGuard,
			blogcontroller.DeleteBlog)
}