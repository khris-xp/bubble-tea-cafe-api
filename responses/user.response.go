package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/types"
)

func AuthUserSuccessResponse(c *fiber.Ctx, statusCode int, message string, token string) error {
	return c.Status(statusCode).JSON(types.AuthResponse{
		Status:  statusCode,
		Message: message,
		Token:   token,
	})
}

func GetUserSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.UserResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func UserErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(types.UserResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}
