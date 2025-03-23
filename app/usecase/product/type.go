package product

import (
	"context"

	"github.com/inventory-service/app/domain/item"
	"github.com/inventory-service/app/domain/product"
	"github.com/inventory-service/app/dto"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

type ProductService interface {
	Create(ctx context.Context, name string, ingredients []dto.Ingredient) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, prodcutID string) (model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
}

type productService struct {
	productDomain product.ProductDomain
	itemDomain    item.ItemDomain
}
