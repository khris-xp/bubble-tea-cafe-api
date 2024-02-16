package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/controllers"
	"github.com/khris-xp/bubble-milk-tea/middlewares"
	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func OrderRoutes(app *fiber.App) {
	orderRepo := repositories.NewOrderRepository()
	orderController := controllers.NewOrderController(orderRepo)

	order := app.Group("/api/order")
	order.Get("/", middlewares.AuthStaffMiddleware(), orderController.GetAllOrder)
	order.Get("/:id", middlewares.AuthMiddleware(), orderController.GetOrderByID)
	order.Get("/user/:id", middlewares.AuthMiddleware(), orderController.GetOrderByUserId)
	order.Post("/", middlewares.AuthMiddleware(), orderController.CreateOrder)
	order.Put("/status/:id", middlewares.AuthStaffMiddleware(), orderController.UpdateStatusOrder)
	order.Put("/:id", middlewares.AuthStaffMiddleware(), orderController.UpdateOrder)
	order.Delete("/:id", middlewares.AuthStaffMiddleware(), orderController.DeleteOrder)
}
