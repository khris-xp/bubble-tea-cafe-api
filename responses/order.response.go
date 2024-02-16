package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/types"
)

func OrderSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.OrderResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func OrderErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(types.OrderResponse{
		Status:  statusCode,
		Message: err.Error(),
		Data:    nil,
	})
}

func OrderCreatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.OrderResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func OrderUpdatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.OrderResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func OrderDeletedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.OrderResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}
