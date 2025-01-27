package product

import (
	"context"

	"github.com/inventory-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Create(ctx context.Context, name string, ingredients []model.Ingredient) error
	FindAll(ctx context.Context) ([]model.Product, error)
	FindByID(ctx context.Context, productID string) (model.Product, error)
	Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) error
	Delete(ctx context.Context, productID string) error
}

type productRepository struct {
	productCollection *mongo.Collection
}
