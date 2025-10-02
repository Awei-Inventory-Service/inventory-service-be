package product

import (
	"context"

	inventorydomain "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	inventory "github.com/inventory-service/resource/inventory"

	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
	productrecipe "github.com/inventory-service/resource/product_recipe"
)

type ProductDomain interface {
	Create(ctx context.Context, payload model.Product) (*model.Product, *error_wrapper.ErrorWrapper)
	FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, product dto.UpdateProductRequest, productID string) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
	CalculateProductCost(ctx context.Context, productCompositions []model.ProductRecipe, branchID string) (float64, *error_wrapper.ErrorWrapper)
}

type productDomain struct {
	inventoryDomain       inventorydomain.InventoryDomain
	productResource       product.ProductResource
	itemResource          item.ItemResource
	productRecipeResource productrecipe.ProductRecipeResource
	inventoryResource     inventory.InventoryResource
}
