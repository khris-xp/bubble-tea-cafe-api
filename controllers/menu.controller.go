package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuController struct {
	MenuRepo *repositories.MenuRepository
}

func NewMenuController(menuRepo *repositories.MenuRepository) *MenuController {
	return &MenuController{MenuRepo: menuRepo}
}

func (mc *MenuController) GetAllMenu(c *fiber.Ctx) error {
	menu, err := mc.MenuRepo.GetAllMenu(c.Context())
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.MenuSuccessResponse(c, fiber.StatusOK, "success", menu)
}

func (mc *MenuController) GetMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")
	menu, err := mc.MenuRepo.GetMenuByID(c.Context(), id)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return responses.MenuSuccessResponse(c, fiber.StatusOK, "success", menu)
}

func (mc *MenuController) CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	menu, err := mc.MenuRepo.CreateMenu(c.Context(), menu)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.CreateMenuSuccessResponse(c, fiber.StatusCreated, "success", menu)
}

func (mc *MenuController) UpdateMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	menuID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid menu id")
	}

	menu, err = mc.MenuRepo.UpdateMenu(c.Context(), menuID, menu)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.UpdateMenuSuccessResponse(c, fiber.StatusOK, "success", menu)
}

func (mc *MenuController) DeleteMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	menuID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid menu id")
	}

	err = mc.MenuRepo.DeleteMenu(c.Context(), menuID)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.DeleteMenuSuccessResponse(c, fiber.StatusOK, "success", nil)
}

func (mc *MenuController) AddToppingToMenu(c *fiber.Ctx) error {
	var topping models.Topping
	id := c.Params("id")
	if err := c.BodyParser(&topping); err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	menuID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid menu id")
	}
	toppingID, err := primitive.ObjectIDFromHex(topping.Id)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusBadRequest, "invalid menu id")
	}

	err = mc.MenuRepo.AddToppingToMenu(c.Context(), menuID, toppingID)
	if err != nil {
		return responses.MenuErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.AddToppingToMenuSuccessResponse(c, fiber.StatusCreated, "success", topping)
}
