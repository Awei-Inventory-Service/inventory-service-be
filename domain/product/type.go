package product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
)

type ProductDomain interface {
	Create(ctx context.Context, payload model.Product) (*model.Product, *error_wrapper.ErrorWrapper)
	FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, product model.Product) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
}

type productDomain struct {
	productResource product.ProductResource
	itemResource    item.ItemResource
}
