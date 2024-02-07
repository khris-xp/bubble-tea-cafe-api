package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/controllers"
	"github.com/khris-xp/bubble-milk-tea/middlewares"
	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func CategoryRoutes(app *fiber.App) {
	categoryRepo := repositories.NewCategoryRepository()
	categoryController := controllers.NewCategoryController(categoryRepo)

	category := app.Group("/api/category")
	category.Get("/", categoryController.GetAllCategories)
	category.Get("/:id", categoryController.GetCategoryByID)
	category.Post("/", middlewares.AuthStaffMiddleware(), categoryController.CreateCategory)
	category.Put("/:id", middlewares.AuthStaffMiddleware(), categoryController.UpdateCategory)
	category.Delete("/:id", middlewares.AuthStaffMiddleware(), categoryController.DeleteCategory)
}
