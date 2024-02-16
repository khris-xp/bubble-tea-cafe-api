package controllers

import (
	"time"

	"github.com/khris-xp/bubble-milk-tea/configs"
	"github.com/khris-xp/bubble-milk-tea/models"
	"github.com/khris-xp/bubble-milk-tea/repositories"
	"github.com/khris-xp/bubble-milk-tea/responses"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
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

func (ac *AuthController) AddMenuToCart(c *fiber.Ctx) error {
	var cart models.Cart
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

	if err := c.BodyParser(&cart); err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	for _, item := range user.Cart {
		if item.MenuId == cart.MenuId {
			return responses.UserErrorResponse(c, fiber.StatusBadRequest, "menu already exists in the cart")
		}
	}

	cart.UserId = user.Id
	cart.Id = primitive.NewObjectID().Hex()
	userID, err := primitive.ObjectIDFromHex(user.Id)
	cart.Status = "pending"
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()

	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid user id")
	}
	_, err = ac.UserRepo.AddMenuToCart(c.Context(), cart, userID)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return responses.AddMenuToCartSuccessResponse(c, fiber.StatusCreated, "success", cart)
}

func (ac *AuthController) EditMenuInCart(c *fiber.Ctx) error {
	cartId := c.Params("id")
	var cart models.Cart
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

	if err := c.BodyParser(&cart); err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid cart id")
	}

	userID, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid user id")
	}

	cart.Id = cartId
	cart.UserId = user.Id
	cart.Status = "pending"
	cart.UpdatedAt = time.Now()

	_, err = ac.UserRepo.EditMenuInCart(c.Context(), cart, userID, id)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return responses.EditMenuInCartSuccessResponse(c, fiber.StatusOK, "success", cart)

}

func (ac *AuthController) RemoveMenuFromCart(c *fiber.Ctx) error {
	cartId := c.Params("id")

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

	id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid cart id")
	}

	user_id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid user id")
	}

	_, err = ac.UserRepo.RemoveMenuFromCart(c.Context(), id, user_id)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.RemoveMenuFromCartSuccessResponse(c, fiber.StatusOK, "success", nil)
}

func (ac *AuthController) RemoveAllMenuinCart(c *fiber.Ctx) error {
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

	user_id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusBadRequest, "invalid user id")
	}

	_, err = ac.UserRepo.RemoveAllMenuinCart(c.Context(), user_id)
	if err != nil {
		return responses.UserErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.RemoveMenuFromCartSuccessResponse(c, fiber.StatusOK, "success", nil)
}
