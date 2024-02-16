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

var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "order")

type OrderRepositoryInterface interface {
	GetAllOrder(ctx context.Context) ([]models.Order, error)
	GetOrderByID(ctx context.Context, id string) (models.Order, error)
	CreateOrder(ctx context.Context, order models.Order) (models.Order, error)
	UpdateOrder(ctx context.Context, id string, order models.Order) (models.Order, error)
	DeleteOrder(ctx context.Context, id string) (int64, error)
}

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{collection: orderCollection}
}

func (or *OrderRepository) GetAllOrder(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := orderCollection.Find(ctx, bson.D{})
	if err != nil {
		return orders, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderByID(ctx context.Context, id string) (models.Order, error) {
	var order models.Order
	orderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = orderCollection.FindOne(ctx, bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (or *OrderRepository) GetOrderByUserId(ctx context.Context, userId string) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := orderCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return orders, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		return order, err
	}

	order.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return order, nil
}

func (or *OrderRepository) UpdateOrder(ctx context.Context, id string, order models.Order) (models.Order, error) {
	order.UpdatedAt = time.Now()

	orderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	_, err = orderCollection.UpdateOne(ctx, bson.M{"_id": orderID}, bson.M{"$set": order})
	if err != nil {
		return order, err
	}

	return order, nil
}

func (or *OrderRepository) UpdateStatusOrder(ctx context.Context, id string, status string) (models.Order, error) {
	var order models.Order
	orderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = orderCollection.FindOne(ctx, bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		return order, err
	}

	order.Status = status
	order.UpdatedAt = time.Now()

	_, err = orderCollection.UpdateOne(ctx, bson.M{"_id": orderID}, bson.M{"$set": order})
	if err != nil {
		return order, err
	}

	return order, nil
}

func (or *OrderRepository) DeleteOrder(ctx context.Context, id string) (int64, error) {
	orderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": orderID})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
