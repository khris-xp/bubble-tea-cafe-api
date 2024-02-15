package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToppingController struct {
	ToppingRepo *repositories.ToppingRepository
}

func NewToppingController(toppingRepo *repositories.ToppingRepository) *ToppingController {
	return &ToppingController{ToppingRepo: toppingRepo}
}

func (tc *ToppingController) CreateTopping(c *fiber.Ctx) error {
	var topping models.Topping
	if err := c.BodyParser(&topping); err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	id, err := tc.ToppingRepo.CreateTopping(c.Context(), topping)
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return responses.ToppingCreatedResponse(c, fiber.StatusCreated, id)
}

func (tc *ToppingController) GetToppingByID(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	topping, err := tc.ToppingRepo.GetToppingByID(c.Context(), id)
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return responses.ToppingSuccessResponse(c, fiber.StatusOK, topping)
}

func (tc *ToppingController) GetAllToppings(c *fiber.Ctx) error {
	var toppings []models.Topping

	toppings, err := tc.ToppingRepo.GetAllToppings(c.Context())
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return responses.ToppingSuccessResponse(c, fiber.StatusOK, toppings)
}

func (tc *ToppingController) UpdateTopping(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	var topping models.Topping
	if err := c.BodyParser(&topping); err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = tc.ToppingRepo.UpdateTopping(c.Context(), id, topping)
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return responses.ToppingUpdatedResponse(c, fiber.StatusOK, id)
}

func (tc *ToppingController) DeleteTopping(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = tc.ToppingRepo.DeleteTopping(c.Context(), id)
	if err != nil {
		return responses.ToppingErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return responses.ToppingDeletedResponse(c, fiber.StatusOK, id)
}
