package inventorystockcount

import (
	"context"
	"time"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *inventoryStockCountDomain) Create(ctx context.Context, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper {
	newData := model.InventoryStockCount{
		BranchID:  branchID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Items:     items,
	}
	return i.inventoryStockCountResource.Create(ctx, newData)
}

func (i *inventoryStockCountDomain) Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper {
	return i.inventoryStockCountResource.Update(ctx, stockCountID, branchID, items)
}

func (i *inventoryStockCountDomain) FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	return i.inventoryStockCountResource.FindAll(ctx)
}

func (i *inventoryStockCountDomain) FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	return i.inventoryStockCountResource.FindByID(ctx, stockCountID)
}

func (i *inventoryStockCountDomain) FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	return i.inventoryStockCountResource.FilterByBranch(ctx, branchID)
}

func (i *inventoryStockCountDomain) Delete(ctx context.Context, stockCountID string) *error_wrapper.ErrorWrapper {
	return i.inventoryStockCountResource.Delete(ctx, stockCountID)
}
