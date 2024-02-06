package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/khris-xp/bubble-milk-tea/repositories"
)

func AuthStaffMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !parsedToken.Valid {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		email, ok := claims["email"].(string)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		user, err := repositories.NewUserRepository().GetUserProfile(c.Context(), email)
		role := user.Role

		if role != "staff" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized Role",
			})
		}

		c.Locals("email", email)
		c.Locals("role", role)

		return c.Next()
	}
}
