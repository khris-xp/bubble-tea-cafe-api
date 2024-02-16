package repositories

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/khris-xp/bubble-milk-tea/configs"
	"github.com/khris-xp/bubble-milk-tea/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var userTimeout = 10 * time.Second

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, user models.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error)
	AddWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error
	RemoveWishListByMenuID(ctx context.Context, menuID primitive.ObjectID, userID primitive.ObjectID) error
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: configs.GetCollection(configs.DB, "users"),
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	existingUser, err := userCollection.Find(ctx, bson.M{"email": user.Email})

	if existingUser.Next(ctx) {
		return "", err
	} else if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)
	user.Role = "customer"
	user.Cart = []models.Cart{}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepository) LoginUser(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepository) GetUserProfile(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) AddMenuToCart(ctx context.Context, cart models.Cart, userID primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	menuExists := false
	for _, item := range user.Cart {
		if item.MenuId == cart.MenuId {
			menuExists = true
			break
		}
	}

	if menuExists {
		return user, nil
	}

	user.Cart = append(user.Cart, cart)

	_, err = userCollection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"cart": user.Cart}},
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) RemoveMenuFromCart(ctx context.Context, cartID primitive.ObjectID, userId primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	var updatedCart []models.Cart
	for _, item := range user.Cart {
		id, err := primitive.ObjectIDFromHex(item.Id)
		if err != nil {
			continue
		}
		if id != cartID {
			updatedCart = append(updatedCart, item)
		}
	}

	_, err = userCollection.UpdateOne(
		ctx,
		bson.M{"_id": userId},
		bson.M{"$set": bson.M{"cart": updatedCart}},
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
