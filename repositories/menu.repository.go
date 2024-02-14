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

var menuCollection *mongo.Collection = configs.GetCollection(configs.DB, "menu")
var menuTimeOut = 10 * time.Second

type MenuReepositoryInterface interface {
	GetAllMenu(ctx context.Context) ([]models.Menu, error)
	GetMenuByID(ctx context.Context, id primitive.ObjectID) (models.Menu, error)
	CreateMenu(ctx context.Context, menu models.Menu) (models.Menu, error)
	UpdateMenu(ctx context.Context, id primitive.ObjectID, menu models.Menu) (models.Menu, error)
	DeleteMenu(ctx context.Context, id primitive.ObjectID) (int64, error)
}

type MenuRepository struct {
	collection *mongo.Collection
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{collection: menuCollection}
}

func (mr *MenuRepository) GetAllMenu(ctx context.Context) ([]models.Menu, error) {
	var menus []models.Menu
	cursor, err := menuCollection.Find(ctx, bson.D{})
	if err != nil {
		return menus, err
	}

	if err = cursor.All(ctx, &menus); err != nil {
		return menus, err
	}

	return menus, nil
}

func (mr *MenuRepository) GetMenuByID(ctx context.Context, id string) (models.Menu, error) {
	var menu models.Menu
	menuID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return menu, err
	}

	err = menuCollection.FindOne(ctx, bson.M{"_id": menuID}).Decode(&menu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (mr *MenuRepository) CreateMenu(ctx context.Context, menu models.Menu) (models.Menu, error) {
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()

	_, err := menuCollection.InsertOne(ctx, menu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (mr *MenuRepository) UpdateMenu(ctx context.Context, id primitive.ObjectID, menu models.Menu) (models.Menu, error) {
	menu.UpdatedAt = time.Now()

	_, err := menuCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": menu})
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (mr *MenuRepository) DeleteMenu(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, menuTimeOut)
	defer cancel()

	_, err := menuCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
