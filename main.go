package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/configs"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/khris-xp/bubble-milk-tea/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Bubble Tea API"})
	})

	configs.ConnectDB()
	routes.UserRoutes(app)
	routes.CategoryRoutes(app)
	routes.ToppingRoute(app)
	routes.MenuRoutes(app)

	port := configs.EnvPort()
	app.Listen(":" + port)
}
