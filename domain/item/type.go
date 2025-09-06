package item

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	branchitem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
)

type ItemDomain interface {
	Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, id string) (*dto.GetItemsResponse, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemDomain struct {
	itemResource            item.ItemResource
	itemCompositionResource itemcomposition.ItemCompositionResourece
	purchaseResource        purchase.PurchaseResource
	branchItemResource      branchitem.BranchItemResource
}
