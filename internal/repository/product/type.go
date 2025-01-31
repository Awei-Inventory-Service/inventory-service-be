package product

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
)

type ProductRepository interface {
	Create(ctx context.Context, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, productID string) (model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
}

type productRepository struct {
	productCollection mongodb.MongoDBCollection
}
