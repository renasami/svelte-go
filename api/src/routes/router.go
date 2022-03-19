// routes.go
package routes

import (
	"github.com/renasami/svelte-go/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/other", controllers.Other)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
