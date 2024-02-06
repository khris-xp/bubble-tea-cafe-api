package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/controllers"
	"github.com/khris-xp/bubble-milk-tea/middlewares"
	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func UserRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	userController := controllers.NewAuthController(userRepo)

	auth := app.Group("/api/auth")
	auth.Post("/register", userController.Register)
	auth.Post("/login", userController.Login)

	auth.Get("/user/:id", middlewares.AuthMiddleware(), userController.GetUserById)
	auth.Get("/profile", middlewares.AuthMiddleware(), userController.GetUserProfile)
	auth.Get("/users", middlewares.AuthStaffMiddleware(), userController.GetAllUsers)
}
