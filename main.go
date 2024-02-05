package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/configs"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Bubble Tea API"})
	})

	configs.ConnectDB()

	port := configs.EnvPort()
	app.Listen(":" + port)
}
