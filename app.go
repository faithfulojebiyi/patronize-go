package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": "Welcome to patronise"})

	})

	app.Listen(":3002")
}
