package product

import (
	inventory_domain "github.com/inventory-service/domain/inventory"
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/inventory_snapshot"
	"github.com/inventory-service/resource/product_snapshot"

	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
	productrecipe "github.com/inventory-service/resource/product_recipe"
)

func NewProductDomain(
	productResource product.ProductResource,
	itemResource item.ItemResource,
	productRecipeResource productrecipe.ProductRecipeResource,
	inventoryResource inventory.InventoryResource,
	inventoryDomain inventory_domain.InventoryDomain,
	inventorySnapshotResource inventory_snapshot.InventorySnapshotResource,
	productSnapshotResource product_snapshot.ProductSnaspshotResource,
) ProductDomain {
	return &productDomain{
		productResource:           productResource,
		itemResource:              itemResource,
		productRecipeResource:     productRecipeResource,
		inventoryResource:         inventoryResource,
		inventoryDomain:           inventoryDomain,
		inventorySnapshotResource: inventorySnapshotResource,
		productSnapshotResource:   productSnapshotResource,
	}
}
