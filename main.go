package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/configs"

	"github.com/khris-xp/bubble-milk-tea/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Bubble Tea API"})
	})

	configs.ConnectDB()
	routes.UserRoutes(app)
	routes.CategoryRoutes(app)

	port := configs.EnvPort()
	app.Listen(":" + port)
}
