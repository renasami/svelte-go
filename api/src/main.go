// main.go
package main

import (
	"github.com/renasami/svelte-go/api/database"
	"github.com/renasami/svelte-go/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// GORMセット
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8080")
}
