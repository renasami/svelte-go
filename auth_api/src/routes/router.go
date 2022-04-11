package routes

import (
	"github.com/renasami/svelte-go/auth_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/user", controllers.User)
	app.Get("/api/auth/logout", controllers.Logout)
	app.Post("/api/auth/forgot", controllers.Forgot)
	app.Post("/api/auth/reset", controllers.Reset)
}
