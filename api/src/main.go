// main.go
package main

import (
	"github.com/renasami/svelte-go/api/database"
	"github.com/renasami/svelte-go/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// GORMセット
	database.Connect()

	app := fiber.New()
	//corrs
	app.Use(cors.New(cors.Config{
		// https://docs.gofiber.io/api/middleware/cors#config
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
