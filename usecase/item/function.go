package item

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/utils"
	"github.com/inventory-service/model"
)

func (i *itemUsecase) Create(ctx context.Context, payload dto.CreateItemRequest) (errW *error_wrapper.ErrorWrapper) {
	itemCategory, errW := utils.ParseItemCategory(payload.Category)

	if errW != nil {
		return
	}

	compositions := model.ItemCompositions{
		Compositions: make([]model.CompositionItem, len(payload.ItemCompositions)),
	}

	for idx, comp := range payload.ItemCompositions {
		compItem, errW := i.itemDomain.FindByID(ctx, comp.ItemID)
		if errW != nil {
			return errW
		}

		compositions.Compositions[idx] = model.CompositionItem{
			ItemID:   comp.ItemID,
			ItemName: compItem.Name,
			Unit:     compItem.Unit,
		}
	}

	_, errW = i.itemDomain.Create(model.Item{
		Name:         payload.Name,
		Category:     itemCategory,
		Compositions: compositions,
		Unit:         payload.Unit,
		SupplierID: func() *string {
			if payload.SupplierID == "" {
				return nil
			}
			return &payload.SupplierID
		}(),
	})

	if errW != nil {
		return errW
	}

	return nil
}

func (i *itemUsecase) FindAll() ([]dto.GetItemsResponse, *error_wrapper.ErrorWrapper) {
	items, err := i.itemDomain.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemUsecase) FindByID(ctx context.Context, id string) (*dto.GetItemsResponse, *error_wrapper.ErrorWrapper) {
	item, err := i.itemDomain.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemUsecase) Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) (errW *error_wrapper.ErrorWrapper) {

	errW = i.itemDomain.Update(ctx, payload, itemID)
	if errW != nil {
		return errW
	}

	return nil
}

func (i *itemUsecase) Delete(id string) *error_wrapper.ErrorWrapper {
	err := i.itemDomain.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
