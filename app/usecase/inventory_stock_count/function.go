package inventorystockcount

import (
	"context"

	"github.com/inventory-service/app/dto"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *inventoryStockCountService) Create(ctx context.Context, branchID string, items []dto.StockCount) *error_wrapper.ErrorWrapper {
	branch, err := i.branchDomain.FindByID(branchID)

	if err != nil {
		return err
	}

	if branch.UUID == "" {
		return error_wrapper.New(model.SErrBranchNotExist, "Branch doesn't exist").With(branchID)
	}
	var itemsData []model.ItemCount
	for _, item := range items {
		itemData, err := i.itemDomain.FindByID(item.ItemID)
		if err != nil {
			return err
		}

		if itemData.UUID == "" {
			return error_wrapper.New(model.SErrItemNotExist, "Item doesn't exist").With(itemData.UUID)
		}

		itemsData = append(itemsData, model.ItemCount{
			ItemName:     itemData.Name,
			ItemID:       itemData.UUID,
			CurrentStock: item.CurrentStock,
		})

	}

	err = i.inventoryStockCountDomain.Create(ctx, branchID, itemsData)

	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryStockCountService) Update(ctx context.Context, stockCountID string, branchID string, items []dto.StockCount) *error_wrapper.ErrorWrapper {
	branch, err := i.branchDomain.FindByID(branchID)

	if err != nil {
		return err
	}

	if branch.UUID == "" {
		return error_wrapper.New(model.SErrBranchNotExist, "Branch doesn't exist").With(branchID)
	}
	var itemsData []model.ItemCount
	for _, item := range items {
		itemData, err := i.itemDomain.FindByID(item.ItemID)
		if err != nil {
			return err
		}

		if itemData.UUID == "" {
			return error_wrapper.New(model.SErrItemNotExist, "Item doesn't exist").With(itemData.UUID)
		}

		itemsData = append(itemsData, model.ItemCount{
			ItemName:     itemData.Name,
			ItemID:       itemData.UUID,
			CurrentStock: item.CurrentStock,
		})

	}

	err = i.inventoryStockCountDomain.Update(ctx, stockCountID, branchID, itemsData)

	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryStockCountService) FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	inventoryStockCounts, err := i.inventoryStockCountDomain.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountService) FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	inventoryStockCount, err := i.inventoryStockCountDomain.FindByID(ctx, stockCountID)

	if err != nil {
		return model.InventoryStockCount{}, err
	}

	return inventoryStockCount, nil
}

func (i *inventoryStockCountService) FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	branch, err := i.branchDomain.FindByID(branchID)
	if err != nil {
		return nil, err
	}

	if branch.UUID == "" {
		return nil, error_wrapper.New(model.SErrBranchNotExist, "Branch doesn't exist").With(branchID)

	}

	inventoryStockCounts, err := i.inventoryStockCountDomain.FilterByBranch(ctx, branchID)

	if err != nil {
		return nil, err
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountService) Delete(ctx context.Context, stockCountID string) *error_wrapper.ErrorWrapper {

	return i.inventoryStockCountDomain.Delete(ctx, stockCountID)
}
