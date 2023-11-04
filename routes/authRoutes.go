package routes

import (
	authusercontroller "github.com/Dandy-CP/go-fiber-portfolio/controllers/AuthUserController"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(auth fiber.Router) {
	auth.Post("/login",
		middleware.ValidateAuth,
			authusercontroller.AuthLogin)

	auth.Post("/register",
		middleware.ValidateAuth,
			authusercontroller.AuthSignUp)

	auth.Get("/logout",
		middleware.AuthGuard,
			authusercontroller.LogoutUser)
}