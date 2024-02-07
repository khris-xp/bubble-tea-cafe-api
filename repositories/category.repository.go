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

var categoryCollection *mongo.Collection = configs.GetCollection(configs.DB, "categories")
var categoryTimeout = 10 * time.Second

type CategoryRepositoryInterface interface {
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id primitive.ObjectID) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, id primitive.ObjectID, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
}

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		collection: configs.GetCollection(configs.DB, "categories"),
	}
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	cursor, err := categoryCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id primitive.ObjectID) (models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	var category models.Category
	err := categoryCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepository) GetCategoryByName(ctx context.Context, name string) (models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	var category models.Category
	err := categoryCollection.FindOne(ctx, bson.M{"name": name}).Decode(&category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	_, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id primitive.ObjectID, category models.Category) (models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	category.UpdatedAt = time.Now()

	_, err := categoryCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": category})
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, categoryTimeout)
	defer cancel()

	_, err := categoryCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
