package product

import (
	"context"

	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
	productcomposition "github.com/inventory-service/domain/product_composition"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ProductService interface {
	Create(ctx context.Context, payload dto.CreateProductRequest) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, prodcutID string) (*model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload model.Product) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
}

type productService struct {
	productDomain            product.ProductDomain
	itemDomain               item.ItemDomain
	productCompositionDomain productcomposition.ProductCompositionDomain
}
