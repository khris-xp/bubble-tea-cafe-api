package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/types"
)

func MenuSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	menuResponse := types.MenuResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(menuResponse)
}

func MenuErrorResponse(c *fiber.Ctx, status int, message string) error {
	menuResponse := types.MenuResponse{
		Status:  status,
		Message: message,
		Data:    nil,
	}

	return c.Status(status).JSON(menuResponse)
}

func CreateMenuSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	menuResponse := types.MenuResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(menuResponse)
}

func UpdateMenuSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	menuResponse := types.MenuResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(menuResponse)
}

func DeleteMenuSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	menuResponse := types.MenuResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(menuResponse)
}
