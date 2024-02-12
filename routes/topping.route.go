package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/controllers"
	"github.com/khris-xp/bubble-milk-tea/middlewares"
	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func ToppingRoute(app *fiber.App) {
	toppingRepo := repositories.NewToppingRepository()
	toppingController := controllers.NewToppingController(toppingRepo)

	topping := app.Group("/api/topping")
	topping.Get("/", toppingController.GetAllToppings)
	topping.Get("/:id", toppingController.GetToppingByID)
	topping.Post("/", middlewares.AuthStaffMiddleware(), toppingController.CreateTopping)
	topping.Put("/:id", middlewares.AuthStaffMiddleware(), toppingController.UpdateTopping)
	topping.Delete("/:id", middlewares.AuthStaffMiddleware(), toppingController.DeleteTopping)
}
