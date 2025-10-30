package product

import (
	"context"
	"time"

	inventorydomain "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/inventory_snapshot"
	"github.com/inventory-service/resource/product_snapshot"

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
	CalculateProductCost(ctx context.Context, product model.Product, branchID string, timestamp time.Time) ([]dto.ProductRecipeWithPrice, float64, *error_wrapper.ErrorWrapper)
}

type productDomain struct {
	inventoryDomain           inventorydomain.InventoryDomain
	productResource           product.ProductResource
	itemResource              item.ItemResource
	productRecipeResource     productrecipe.ProductRecipeResource
	inventoryResource         inventory.InventoryResource
	inventorySnapshotResource inventory_snapshot.InventorySnapshotResource
	productSnapshotResource   product_snapshot.ProductSnaspshotResource
}
