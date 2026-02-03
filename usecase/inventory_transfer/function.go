package inventory_transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/constant"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (i *inventoryTransferUsecase) Create(ctx context.Context, payload dto.CreateInventoryTransferRequest) (newData model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	newData, errW = i.inventoryTransferDomain.Create(ctx, payload)
	if errW != nil {
		fmt.Println("Error creating inventory transfer", errW)
		return
	}

	parsedTime, err := time.Parse("2006-01-02", payload.TransferDate)
	if err != nil {
		return newData, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
	}

	referenceTypeInventoryTransfer := constant.InventoryTransfer
	for _, transferItem := range payload.Items {
		// Sync brnach item in destination id, so that it get the correct price
		_, _, _ = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchOriginID, transferItem.ItemID)
		inventory, errW := i.inventoryDomain.GetInventoryByDate(ctx, parsedTime, transferItem.ItemID, payload.BranchOriginID)

		if errW != nil {
			fmt.Println("Error getting inventory price", errW)
			continue
		}

		item, errW := i.itemDomain.FindByID(ctx, transferItem.ItemID)
		if errW != nil {
			fmt.Println("Error getting item by id in create inventory transfer")
			return newData, errW
		}

		standarizeUnit := utils.StandarizeMeasurement(transferItem.Quantity, transferItem.Unit, item.Unit)
		fmt.Printf("Inventory price: %f and standarize unit: %f", inventory.Price, standarizeUnit)
		_, errW = i.inventoryTransferItemDomain.Create(ctx, model.InventoryTransferItem{
			InventoryTransferID: newData.UUID,
			ItemID:              transferItem.ItemID,
			ItemQuantity:        transferItem.Quantity,
			Unit:                transferItem.Unit,
			ItemCost:            inventory.Price * standarizeUnit,
		})
		if errW != nil {
			fmt.Println("Error creating inventory transfer", errW)
			continue
		}

		parsedTime, err := time.Parse("2006-01-02", payload.TransferDate)
		if err != nil {
			return model.InventoryTransfer{}, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
		}

		errW = i.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchOriginID,
			BranchDestinationID: payload.BranchDestinationID,
			ItemID:              item.UUID,
			Type:                "OUT",
			Quantity:            transferItem.Quantity,
			IssuerID:            payload.IssuerID,
			Unit:                transferItem.Unit,
			Reference:           newData.UUID,
			Cost:                inventory.Price * standarizeUnit,
			ReferenceType:       &referenceTypeInventoryTransfer,
			TransactionDate:     parsedTime,
		})
		if errW != nil {
			fmt.Println("Error creating new stock transaction", errW)
			return model.InventoryTransfer{}, errW
		}

		_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchOriginID, transferItem.ItemID)
		if errW != nil {
			fmt.Println("Error sync branch item", errW, transferItem.ItemID)
			continue
		}

		errW = i.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			ItemID:   transferItem.ItemID,
			BranchID: payload.BranchOriginID,
			NewTime:  payload.TransferDate,
		})
		if errW != nil {
			fmt.Println("Error recalculating inventory domain", errW)
			continue
		}
	}
	return
}

func (i *inventoryTransferUsecase) UpdateStatus(ctx context.Context, payload dto.UpdateInventoryTransferStatus) (errW *error_wrapper.ErrorWrapper) {
	statusValid := payload.ValidateStatus()

	if !statusValid {
		errW = error_wrapper.New(model.UErrInvalidInventoryTransferStatus, "Status not valid")
		return
	}

	errW = i.inventoryTransferDomain.UpdateStatus(ctx, payload.InventoryTransferID, payload.Status)
	if errW != nil {
		fmt.Println("Error updating inventory transfer status", errW)
		return
	}

	inventoryTransfer, errW := i.inventoryTransferDomain.FindByID(ctx, payload.InventoryTransferID)
	if errW != nil {
		fmt.Println("Error getting inventory transfer based on id ", errW)
		return
	}

	if payload.Status == constant.TRANSFER_STATUS_COMPLETED {

		// If an inventory status is completed, then create new stock transaction
		for _, inventoryTransferItem := range inventoryTransfer.Items {
			referenceType := constant.InventoryTransfer
			stockTransaction := model.StockTransaction{
				BranchOriginID:      inventoryTransfer.BranchOriginID,
				BranchDestinationID: inventoryTransfer.BranchDestinationID,
				ItemID:              inventoryTransferItem.ItemID,
				Type:                "IN",
				Quantity:            inventoryTransferItem.ItemQuantity,
				IssuerID:            inventoryTransfer.IssuerID,
				Unit:                inventoryTransferItem.Unit,
				Reference:           inventoryTransfer.UUID,
				Cost:                inventoryTransferItem.ItemCost,
				ReferenceType:       &referenceType,
				TransactionDate:     time.Now(),
			}
			errW = i.stockTransactionDomain.Create(stockTransaction)

			if errW != nil {
				fmt.Println("Error creating stock transaction for IN process in UpdateStatus", errW)
				return
			}
			_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, inventoryTransfer.BranchDestinationID, inventoryTransferItem.ItemID)
			if errW != nil {
				fmt.Println("Error sync branch item", errW, inventoryTransferItem.ItemID)
				continue
			}
			errW := i.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
				BranchID: inventoryTransfer.BranchDestinationID,
				ItemID:   inventoryTransferItem.ItemID,
				NewTime:  time.Now().Format("2006-01-02"),
			})
			if errW != nil {
				fmt.Println("Error recalculating inventory", errW)
			}

			// _, _, errW = i.inventoryDomain.SyncBranchItem(ctx, inventoryTransfer.BranchOriginID, inventoryTransferItem.ItemID)
			// if errW != nil {
			// 	fmt.Println("Error sync branch item", errW, inventoryTransferItem.ItemID)
			// 	continue
			// }

		}
	}
	return
}

func (i *inventoryTransferUsecase) Get(ctx context.Context, payload dto.GetListRequest) (result dto.GetInventoryTransferListResponse, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryTransferDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
}

func (i *inventoryTransferUsecase) Update(ctx context.Context, id string, payload dto.UpdateInventoryTransferRequest) (result model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	var (
		updatedItems []string
	)

	oldInventoryTransfer, errW := i.inventoryTransferDomain.FindByID(ctx, id)
	if errW != nil {
		fmt.Println("Error finding old inventory transfer by id ", id)
		return
	}

	parsedTransferDate, err := time.Parse("2006-01-02", payload.TransferDate)
	if err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, "Invalid date format")
		fmt.Println("ERROR", errW)
		return
	}
	// 1. Update the old inventory transfer
	updatedInventoryTransfer, errW := i.inventoryTransferDomain.Update(ctx, id, payload)
	if errW != nil {
		fmt.Println("Error updating inventory transfer", errW)
		return
	}

	// 2. Delete the old inventory transfer item and the stock transaction
	errW = i.inventoryTransferItemDomain.Delete(ctx, model.InventoryTransferItem{
		InventoryTransferID: id,
	})

	if errW != nil {
		fmt.Println("Error deleting inventory transfer item", errW)
		return
	}

	items, errW := i.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": id,
		},
	}, payload.IssuerID)

	if errW != nil {
		fmt.Println("Fail invalidating stock transaction")
		return updatedInventoryTransfer, errW
	}
	updatedItems = append(updatedItems, items...)

	now := time.Now()
	// 3. Create new inventory transfer item and stock transaction
	for _, transferItem := range payload.Items {
		_, _, _ = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchOriginID, transferItem.ItemID)
		inventory, errW := i.inventoryDomain.GetInventoryByDate(ctx, now, transferItem.ItemID, payload.BranchOriginID)

		if errW != nil {
			fmt.Println("Error getting inventory price", errW)
			continue
		}

		item, errW := i.itemDomain.FindByID(ctx, transferItem.ItemID)

		if errW != nil {
			fmt.Println("Error getting item by id in create inventory transfer")
			return updatedInventoryTransfer, errW
		}

		standarizeUnit := utils.StandarizeMeasurement(transferItem.Quantity, transferItem.Unit, item.Unit)

		_, errW = i.inventoryTransferItemDomain.Create(ctx, model.InventoryTransferItem{
			InventoryTransferID: id,
			ItemID:              transferItem.ItemID,
			ItemQuantity:        transferItem.Quantity,
			Unit:                transferItem.Unit,
			ItemCost:            inventory.Price * standarizeUnit,
		})

		if errW != nil {
			fmt.Println("Error creating inventory transfer", errW)
			continue
		}

		referenceType := constant.InventoryTransfer
		stockTransactionOut := model.StockTransaction{
			BranchOriginID:      updatedInventoryTransfer.BranchOriginID,
			BranchDestinationID: updatedInventoryTransfer.BranchDestinationID,
			ItemID:              transferItem.ItemID,
			Type:                "OUT",
			Quantity:            transferItem.Quantity,
			IssuerID:            updatedInventoryTransfer.IssuerID,
			Unit:                transferItem.Unit,
			Reference:           id,
			Cost:                inventory.Price * standarizeUnit,
			ReferenceType:       &referenceType,
			TransactionDate:     parsedTransferDate,
		}
		errW = i.stockTransactionDomain.Create(stockTransactionOut)
		if errW != nil {
			fmt.Println("Error creating stock transaction for OUT process in UpdateStatus", errW)
			return updatedInventoryTransfer, errW
		}

		updatedItems = append(updatedItems, transferItem.ItemID)
		if payload.Status == constant.TRANSFER_STATUS_COMPLETED {
			// IF status is completed, then crate stock transaction

			parsedDate, err := time.Parse("2006-01-02", payload.CompletedDate)
			if err != nil {
				errW = error_wrapper.New(model.CErrJsonBind, "Invalid date format")
				continue
			}

			stockTransactionIn := stockTransactionOut
			stockTransactionIn.Type = "IN"
			stockTransactionIn.TransactionDate = parsedDate

			errW = i.stockTransactionDomain.Create(stockTransactionIn)
			if errW != nil {
				fmt.Println("Error creating stock transaction for IN process in UpdateStatus", errW)
				return updatedInventoryTransfer, errW
			}
			updatedItems = append(updatedItems, transferItem.ItemID)
		}
	}

	var (
		needToBeUpdatedBranch = make(map[string]bool)
	)

	needToBeUpdatedBranch[payload.BranchDestinationID] = true
	needToBeUpdatedBranch[payload.BranchOriginID] = true

	if _, exist := needToBeUpdatedBranch[oldInventoryTransfer.BranchDestinationID]; !exist {
		needToBeUpdatedBranch[oldInventoryTransfer.BranchDestinationID] = true
	}

	if _, exist := needToBeUpdatedBranch[oldInventoryTransfer.BranchOriginID]; !exist {
		needToBeUpdatedBranch[oldInventoryTransfer.BranchOriginID] = true
	}

	for branchId, _ := range needToBeUpdatedBranch {
		errW = i.inventoryDomain.BulkSyncBranchItems(ctx, branchId, updatedItems)
		if errW != nil {
			fmt.Println("Error bulk sync branch items ", errW)
			continue
		}
	}

	// 4.Sync branch item for all impacted items
	errW = i.inventoryDomain.BulkSyncBranchItems(ctx, payload.BranchDestinationID, updatedItems)
	if errW != nil {
		fmt.Println("Error bulk sync branch items for branch destination id", errW)
		return
	}

	return
}

func (i *inventoryTransferUsecase) Delete(ctx context.Context, payload dto.DeleteInventoryTransferRequest) (errW *error_wrapper.ErrorWrapper) {

	//  1. Get inventory transfer old data
	oldInventoryTransfer, errW := i.inventoryTransferDomain.FindByID(ctx, payload.ID)
	if errW != nil {
		fmt.Println("Error getting inventory transfer by id ", payload.ID)
		return
	}

	// 1. Delete inventory transfer
	errW = i.inventoryTransferDomain.Delete(ctx, payload.ID)
	if errW != nil {
		fmt.Println("Error deleting inventory transfer", errW)
		return
	}

	// 2. Delete inventory transfer item
	errW = i.inventoryTransferItemDomain.Delete(ctx, model.InventoryTransferItem{
		InventoryTransferID: payload.ID,
	})

	if errW != nil {
		fmt.Println("Error deleting invnetory transfer item", errW)
		return
	}

	// 3. Invalidate Stock Transaction
	items, errW := i.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": payload.ID,
		},
	}, payload.UserID)

	if errW != nil {
		fmt.Println("Error invalidating stock transaction", errW)
		return
	}

	// 4. Sync branch item
	errW = i.inventoryDomain.BulkSyncBranchItems(ctx, oldInventoryTransfer.BranchOriginID, items)
	if errW != nil {
		fmt.Println("Error bulk sync branch item", errW)
		return
	}

	errW = i.inventoryDomain.BulkSyncBranchItems(ctx, oldInventoryTransfer.BranchDestinationID, items)
	if errW != nil {
		fmt.Println("Error bulk sync branch items for branch destination id ", errW)
		return
	}

	return
}
