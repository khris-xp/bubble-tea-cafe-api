package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/types"
)

func ToppingSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.ToppingResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func ToppingErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(types.ToppingResponse{
		Status:  statusCode,
		Message: err.Error(),
		Data:    nil,
	})
}

func ToppingCreatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.ToppingResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func ToppingUpdatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.ToppingResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func ToppingDeletedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.ToppingResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}
