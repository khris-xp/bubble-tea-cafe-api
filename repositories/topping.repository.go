package repositories

import (
	"context"
	"time"

	"github.com/khris-xp/bubble-milk-tea/configs"
	"github.com/khris-xp/bubble-milk-tea/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var toppingCollection = configs.GetCollection(configs.DB, "toppings")
var toppingTimeout = 10 * time.Second

type ToppingRepositoryInterface interface {
	CreateTopping(ctx context.Context, topping models.Topping) (string, error)
	GetToppingByID(ctx context.Context, id primitive.ObjectID) (models.Topping, error)
	GetAllToppings(ctx context.Context) ([]models.Topping, error)
	UpdateTopping(ctx context.Context, id primitive.ObjectID, topping models.Topping) error
	DeleteTopping(ctx context.Context, id primitive.ObjectID) error
}

type ToppingRepository struct {
	collection *mongo.Collection
}

func NewToppingRepository() *ToppingRepository {
	return &ToppingRepository{
		collection: configs.GetCollection(configs.DB, "toppings"),
	}
}

func (r *ToppingRepository) CreateTopping(ctx context.Context, topping models.Topping) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, toppingTimeout)
	defer cancel()

	topping.CreatedAt = time.Now()
	topping.UpdatedAt = time.Now()

	result, err := toppingCollection.InsertOne(ctx, topping)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *ToppingRepository) GetToppingByID(ctx context.Context, id primitive.ObjectID) (models.Topping, error) {
	ctx, cancel := context.WithTimeout(ctx, toppingTimeout)
	defer cancel()

	var topping models.Topping
	err := toppingCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&topping)
	if err != nil {
		return models.Topping{}, err
	}

	return topping, nil
}

func (r *ToppingRepository) GetAllToppings(ctx context.Context) ([]models.Topping, error) {
	ctx, cancel := context.WithTimeout(ctx, toppingTimeout)
	defer cancel()

	cursor, err := toppingCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var toppings []models.Topping
	err = cursor.All(ctx, &toppings)
	if err != nil {
		return nil, err
	}

	return toppings, nil
}

func (r *ToppingRepository) UpdateTopping(ctx context.Context, id primitive.ObjectID, topping models.Topping) error {
	ctx, cancel := context.WithTimeout(ctx, toppingTimeout)
	defer cancel()

	topping.UpdatedAt = time.Now()
	_, err := toppingCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": topping})
	if err != nil {
		return err
	}

	return nil
}

func (r *ToppingRepository) DeleteTopping(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, toppingTimeout)
	defer cancel()

	_, err := toppingCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
