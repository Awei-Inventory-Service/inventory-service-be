package item

import (
	"context"

	"github.com/inventory-service/domain/item"
	itemcomposition "github.com/inventory-service/domain/item_composition"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
)

type ItemUsecase interface {
	Create(ctx context.Context, payload dto.CreateItemRequest) *error_wrapper.ErrorWrapper
	FindAll() ([]dto.GetItemsResponse, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, id string) (*dto.GetItemsResponse, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemUsecase struct {
	itemDomain            item.ItemDomain
	itemCompositionDomain itemcomposition.ItemCompositionDomain
}
