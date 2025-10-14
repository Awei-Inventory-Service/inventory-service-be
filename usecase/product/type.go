package product

import (
	"context"

	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
	productrecipe "github.com/inventory-service/domain/product_recipe"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ProductService interface {
	Create(ctx context.Context, payload dto.CreateProductRequest) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, prodcutID string) (*model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload dto.UpdateProductRequest, productID string) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
	GetProductCost(ctx context.Context, productID, branchID string) (cost float64, errW *error_wrapper.ErrorWrapper)
}

type productService struct {
	productDomain       product.ProductDomain
	itemDomain          item.ItemDomain
	productRecipeDomain productrecipe.ProductRecipeDomain
	branchProductDomain branch_product.BranchProductDomain
}
