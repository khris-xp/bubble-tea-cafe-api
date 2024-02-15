package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryController(categoryRepo *repositories.CategoryRepository) *CategoryController {
	return &CategoryController{CategoryRepo: categoryRepo}
}

func (cc *CategoryController) GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category

	categories, err := cc.CategoryRepo.GetAllCategories(c.Context())
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategorySuccessResponse(c, fiber.StatusOK, categories)
}

func (cc *CategoryController) GetCategoryByName(c *fiber.Ctx) error {
	name := c.Params("name")

	category, err := cc.CategoryRepo.GetCategoryByName(c.Context(), name)
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategorySuccessResponse(c, fiber.StatusOK, category)
}

func (cc *CategoryController) GetCategoryByID(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	category, err := cc.CategoryRepo.GetCategoryByID(c.Context(), id)
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategorySuccessResponse(c, fiber.StatusOK, category)
}

func (cc *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	category, err := cc.CategoryRepo.CreateCategory(c.Context(), category)
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategoryCreatedResponse(c, fiber.StatusCreated, category)
}

func (cc *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	category, err = cc.CategoryRepo.UpdateCategory(c.Context(), id, category)
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategoryUpdatedResponse(c, fiber.StatusOK, category)
}

func (cc *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	err = cc.CategoryRepo.DeleteCategory(c.Context(), id)
	if err != nil {
		return responses.CategoryErrorResponse(c, fiber.ErrBadRequest.Code, err)
	}

	return responses.CategoryDeletedResponse(c, fiber.StatusOK, nil)
}
