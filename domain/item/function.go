package item

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/utils"
	"github.com/inventory-service/model"
)

func (i *itemDomain) Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.Create(item)
}

func (i *itemDomain) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindAll()
}

func (i *itemDomain) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindByID(id)
}

func (i *itemDomain) Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) *error_wrapper.ErrorWrapper {
	itemCategory, errW := utils.ParseItemCategory(payload.Category)

	if errW != nil {
		return errW
	}

	item := model.Item{
		UUID:       itemID,
		Name:       payload.Name,
		Category:   itemCategory,
		Price:      payload.Price,
		Unit:       payload.Unit,
		SupplierID: &payload.SupplierID,
	}

	updatedItem, errW := i.itemResource.Update(ctx, item)

	if errW != nil {
		return errW
	}

	errW = i.itemCompositionResource.DeleteByItemID(ctx, updatedItem.UUID)

	if errW != nil {
		return errW
	}

	for _, itemComposition := range payload.ItemCompositions {
		errW = i.itemCompositionResource.Create(ctx, model.ItemComposition{
			ParentItemID: updatedItem.UUID,
			ChildItemID:  itemComposition.ItemID,
			Ratio:        itemComposition.Ratio,
			Notes:        itemComposition.Notes,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if errW != nil {
			fmt.Println("Error creating item composition", errW.ActualError())
			return errW
		}
	}
	return nil
}

func (i *itemDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return i.itemResource.Delete(id)
}
