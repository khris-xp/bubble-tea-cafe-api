package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/types"
)

func CategorySuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(types.CategoryResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func CategoryErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(types.CategoryResponse{
		Status:  statusCode,
		Message: err.Error(),
		Data:    nil,
	})
}

func CategoryCreatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(types.CategoryResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func CategoryUpdatedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(types.CategoryResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func CategoryDeletedResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(types.CategoryResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}
