package controllers

import (
	"github.com/khris-xp/bubble-milk-tea/configs"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userValidate = validator.New()
	jwtSecret    = []byte(configs.EnvSecretKey())
)

type AuthController struct {
	UserRepo *repositories.UserRepository
}

func NewAuthController(userRepo *repositories.UserRepository) *AuthController {
	return &AuthController{UserRepo: userRepo}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	tokenString, err := ac.UserRepo.RegisterUser(c.Context(), user)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid email or password")
	} else if err == nil && tokenString == "" {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "email already exists")
	}

	return responses.AuthUserSuccessResponse(c, fiber.StatusCreated, "success", tokenString)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	tokenString, err := ac.UserRepo.LoginUser(c.Context(), user.Email, user.Password)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid email or password")
	}

	return responses.AuthUserSuccessResponse(c, fiber.StatusCreated, "success", tokenString)
}

func (ac *AuthController) GetUserProfile(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return responses.UserErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		return responses.UserErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return responses.UserErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return responses.UserErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	user, err := ac.UserRepo.GetUserProfile(c.Context(), email)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	return responses.GetUserSuccessResponse(c, fiber.StatusOK, user)
}

func (ac *AuthController) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid user id")
	}
	user, err := ac.UserRepo.GetUserById(c.Context(), id)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "user not found")
	}

	return responses.GetUserSuccessResponse(c, fiber.StatusOK, user)
}

func (ac *AuthController) GetAllUsers(c *fiber.Ctx) error {
	users, err := ac.UserRepo.GetAllUsers(c.Context())
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return responses.GetUserSuccessResponse(c, fiber.StatusOK, users)
}
