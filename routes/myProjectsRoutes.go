package routes

import (
	"github.com/Dandy-CP/go-fiber-portfolio/controllers/myprojectscontroller"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/gofiber/fiber/v2"
)

func MyProjectsRoutes(myProjects fiber.Router) {
	myProjects.Get("/", myprojectscontroller.GetProjects)

	myProjects.Post("/",
	middleware.AuthGuard,
		middleware.ValidateMyProjects,
			myprojectscontroller.CreateProjects)

	myProjects.Put("/:id",
	middleware.AuthGuard,
		middleware.ValidateMyProjects,
			myprojectscontroller.UpdateProjects)
	
	myProjects.Delete("/:id",
		middleware.AuthGuard,
			myprojectscontroller.DeleteProjects)
}