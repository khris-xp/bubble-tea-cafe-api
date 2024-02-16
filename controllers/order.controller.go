package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderController(orderRepo *repositories.OrderRepository) *OrderController {
	return &OrderController{OrderRepo: orderRepo}
}

func (oc *OrderController) GetAllOrder(c *fiber.Ctx) error {
	order, err := oc.OrderRepo.GetAllOrder(c.Context())
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return responses.OrderSuccessResponse(c, fiber.StatusOK, order)
}

func (oc *OrderController) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := oc.OrderRepo.GetOrderByID(c.Context(), id)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusNotFound, err)
	}

	return responses.OrderSuccessResponse(c, fiber.StatusOK, order)
}

func (oc *OrderController) GetOrderByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")
	order, err := oc.OrderRepo.GetOrderByUserId(c.Context(), userId)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusNotFound, err)
	}

	return responses.OrderSuccessResponse(c, fiber.StatusOK, order)
}

func (oc *OrderController) CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	order, err := oc.OrderRepo.CreateOrder(c.Context(), order)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return responses.OrderCreatedResponse(c, fiber.StatusCreated, order)
}

func (oc *OrderController) UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	order, err = oc.OrderRepo.UpdateOrder(c.Context(), id, order)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return responses.OrderUpdatedResponse(c, fiber.StatusOK, order)
}

func (or *OrderController) UpdateStatusOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	order, err = or.OrderRepo.UpdateStatusOrder(c.Context(), id, order.Status)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return responses.OrderUpdatedResponse(c, fiber.StatusOK, order)
}

func (oc *OrderController) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusBadRequest, err)
	}

	_, err = oc.OrderRepo.DeleteOrder(c.Context(), id)
	if err != nil {
		return responses.OrderErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return responses.OrderDeletedResponse(c, fiber.StatusOK, "success")
}
