package item

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
	stockbalance "github.com/inventory-service/resource/stock_balance"
)

type ItemDomain interface {
	Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, id string) (*model.Item, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemDomain struct {
	itemResource            item.ItemResource
	itemCompositionResource itemcomposition.ItemCompositionResourece
	purchaseResource        purchase.PurchaseResource
	stockBalanceResource    stockbalance.StockBalanceResource
}
