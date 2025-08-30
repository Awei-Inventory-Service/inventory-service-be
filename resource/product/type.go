package product

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductResource interface {
	Create(ctx context.Context, product model.Product) (*model.Product, *error_wrapper.ErrorWrapper)
	FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, product model.Product) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
}

type productResource struct {
	db *gorm.DB
}
