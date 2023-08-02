package main

import (
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", routes.UserHandler)
	app.Post("/", routes.CreateHandler)

	app.Listen(":3000")
}
