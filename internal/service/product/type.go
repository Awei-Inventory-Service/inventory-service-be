package product

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/product"
)

type ProductService interface {
	Create(ctx context.Context, name string, ingredients []model.Ingredient) error
	FindAll(ctx context.Context) ([]model.Product, error)
	FindByID(ctx context.Context, prodcutID string) (model.Product, error)
	Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) error
	Delete(ctx context.Context, productID string) error
}

type productService struct {
	productRepository product.ProductRepository
}
