package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/controllers"
	"github.com/khris-xp/bubble-milk-tea/middlewares"
	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func MenuRoutes(app *fiber.App) {
	menuRepo := repositories.NewMenuRepository()
	menuController := controllers.NewMenuController(menuRepo)

	menu := app.Group("/api/menu")
	menu.Get("/", menuController.GetAllMenu)
	menu.Get("/:id", menuController.GetMenuByID)
	menu.Post("/", middlewares.AuthStaffMiddleware(), menuController.CreateMenu)
	menu.Put("/:id", middlewares.AuthStaffMiddleware(), menuController.UpdateMenu)
	menu.Delete("/:id", middlewares.AuthStaffMiddleware(), menuController.DeleteMenu)
}
