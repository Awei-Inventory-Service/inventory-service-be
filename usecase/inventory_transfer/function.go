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

	now := time.Now()
	for _, transferItem := range payload.Items {
		// Sync brnach item in destination id, so that it get the correct price
		_, _, _ = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchOriginID, transferItem.ItemID)
		itemCost, errW := i.inventoryDomain.GetPrice(ctx, dto.CustomDate{
			Day:   now.Day(),
			Month: int(now.Month()),
			Year:  now.Year(),
		}, transferItem.ItemID, payload.BranchOriginID)

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

		_, errW = i.inventoryTransferItemDomain.Create(ctx, model.InventoryTransferItem{
			InventoryTransferID: newData.UUID,
			ItemID:              transferItem.ItemID,
			ItemQuantity:        transferItem.Quantity,
			Unit:                transferItem.Unit,
			ItemCost:            itemCost * standarizeUnit,
		})

		if errW != nil {
			fmt.Println("Error creating inventory transfer", errW)
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
			}
			errW = i.stockTransactionDomain.Create(stockTransaction)

			if errW != nil {
				fmt.Println("Error creating stock transaction for IN process in UpdateStatus", errW)
				return
			}

			stockTransactionOut := stockTransaction
			stockTransactionOut.Type = "OUT"

			errW = i.stockTransactionDomain.Create(stockTransactionOut)

			if errW != nil {
				fmt.Println("Error creating stock transaction for OUT process in UpdateStatus", errW)
				return
			}

			_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, inventoryTransfer.BranchDestinationID, inventoryTransferItem.ItemID)

			if errW != nil {
				fmt.Println("Error sync branch item", errW, inventoryTransferItem.ItemID)
				continue
			}

			_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, inventoryTransfer.BranchOriginID, inventoryTransferItem.ItemID)

			if errW != nil {
				fmt.Println("Error sync branch item", errW, inventoryTransferItem.ItemID)
				continue
			}

		}
	}
	return
}

func (i *inventoryTransferUsecase) Get(ctx context.Context, payload dto.GetListRequest) (result dto.GetInventoryTransferListResponse, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryTransferDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
}
