package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (i *inventoryUsecase) Create(ctx context.Context, payload dto.CreateInventoryRequest) *error_wrapper.ErrorWrapper {

	// 1. Create out transaction for the item compositions
	item, errW := i.itemDomain.FindByID(ctx, payload.ItemID)

	if errW != nil {
		return errW
	}
	referenceType := "ITEM_CREATION"

	for _, itemComposition := range item.ChildCompositions {
		fmt.Println("iNI ITEM COMPOSIITON ITEM ID", itemComposition.ChildItemID)
		// total := itemComposition.Ratio * payload.Quantity * itemComposition.PortionSize
		errW := i.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              itemComposition.ChildItemID,
			Type:                "OUT",
			IssuerID:            payload.UserID,
			Quantity:            0,
			Cost:                0.0,
			Unit:                itemComposition.Unit,
			Reference:           "",
			ReferenceType:       &referenceType,
		})

		if errW != nil {
			return errW
		}

		_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, itemComposition.ChildItemID)

		if errW != nil {
			return errW
		}

	}
	// 2. Create the inside transactions for the item

	errW = i.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            payload.UserID,
		Quantity:            payload.Quantity,
		Unit:                item.Unit,
		Reference:           "",
		ReferenceType:       &referenceType,
	})

	if errW != nil {
		return errW
	}

	_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)

	if errW != nil {
		return errW
	}

	return nil
}

func (i *inventoryUsecase) FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindByBranchAndItem(payload.BranchId, payload.ItemId)
}

func (i *inventoryUsecase) FindByBranchId(branchId string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindByBranch(branchId)
}

func (i *inventoryUsecase) GetListCurrent(ctx context.Context, payload dto.GetListRequest, branchID string) (inventories []dto.GetInventoryResponse, errW *error_wrapper.ErrorWrapper) {
	inventories, errW = i.inventoryDomain.Get(ctx, payload)
	if errW != nil {
		fmt.Println("Error getting inventories in get list current usecase", errW)
		return
	}
	fmt.Println("Ini len inventories", len(inventories))
	for _, inventory := range inventories {
		filters := []dto.Filter{
			{
				Key:      "branch_id",
				Values:   []string{branchID},
				Wildcard: "==",
			},
			{
				Key:      "item_id",
				Values:   []string{inventory.ItemID},
				Wildcard: "==",
			},
		}
		_, errW := i.inventorySnapshotDomain.Get(ctx, filters, nil, 0, 0)
		if errW != nil && errW.Is(model.RErrDataNotFound) {
			errW = nil
			// IF data not found, then create inventory snapshot
			errW = i.inventorySnapshotDomain.Upsert(ctx, dto.CreateInventorySnapshotRequest{
				ItemID:   inventory.ItemID,
				BranchID: inventory.BranchID,
				Balance:  inventory.CurrentStock,
				Value:    inventory.Price,
				Date:     time.Now(),
			})
			if errW != nil {
				fmt.Println("Error upserting new inventory snapshot", errW)
				continue
			}
		}
	}
	return
}

func (i *inventoryUsecase) FindAll() ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindAll()
}

func (i *inventoryUsecase) SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (currentStock, currentPrice float64, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)
}

func (i *inventoryUsecase) Get(ctx context.Context, payload dto.GetListRequest, branchID string) ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper) {
	inventorySnapshots, errW := i.inventorySnapshotDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
	fmt.Println("Ini len inventory snapshots", len(inventorySnapshots))

	if dateExist, date := utils.CheckKeyExist("date", payload.Filter); dateExist {
		parsedDate, err := time.Parse("2006-01-02", date[0])
		if err != nil {
			return nil, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
		}

		today := time.Now().Truncate(24 * time.Hour)
		parsedDateTruncated := parsedDate.Truncate(24 * time.Hour)
		if parsedDateTruncated.Equal(today) {
			// Remove date filter since inventory table doesn't have date column
			var filteredPayload dto.GetListRequest
			filteredPayload = payload
			filteredPayload.Filter = make([]dto.Filter, 0)
			
			for _, filter := range payload.Filter {
				if filter.Key != "date" {
					filteredPayload.Filter = append(filteredPayload.Filter, filter)
				}
			}
			
			return i.inventoryDomain.Get(ctx, filteredPayload)
		}
	}

	if itemFilterExist, itemID := utils.CheckKeyExist("item_id", payload.Filter); itemFilterExist {
		// 1. Check if the inventory snapshots length equal to the filter
		if len(inventorySnapshots) == len(itemID) {
			fmt.Println("All items requested exist in the database")
			return i.mapInventorySnapshotToResponse(ctx, inventorySnapshots), nil
		}
		// 2. If not equal, then there are missing data that need to be patched
		missingItems := i.mapExistInventoryToRequest(itemID, inventorySnapshots)
		if len(missingItems) > 0 {
			dateExist, date := utils.CheckKeyExist("date", payload.Filter)
			if dateExist {
				parsedTime, err := time.Parse("2006-01-02", date[0])
				if err != nil {
					return nil, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
				}
				// Patching data for data that not exist
				errW = i.BulkCreate(ctx, missingItems, branchID, parsedTime)
				if errW != nil {
					fmt.Println("Error bulk create missing items", errW)
					return nil, errW
				}
			}
		}

		inventorySnapshots, errW = i.inventorySnapshotDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
		if errW != nil {
			fmt.Println("Error getting inventry snapshots after patching", errW)
			return nil, errW
		}

		return i.mapInventorySnapshotToResponse(ctx, inventorySnapshots), nil
	}
	return i.mapInventorySnapshotToResponse(ctx, inventorySnapshots), errW
}

func (i *inventoryUsecase) mapInventorySnapshotToResponse(
	ctx context.Context,
	inventorySnapshots []model.InventorySnapshot,
) (
	response []dto.GetInventoryResponse,
) {

	if len(inventorySnapshots) <= 0 {
		return
	}

	branch, errW := i.branchDomain.FindByID(inventorySnapshots[0].BranchID)
	if errW != nil {
		fmt.Println("Error finding branch by id", errW)
		return
	}

	for _, inventorySnapshot := range inventorySnapshots {
		item, errW := i.itemDomain.FindByID(ctx, inventorySnapshot.ItemID)
		if errW != nil {
			fmt.Println("Error finding item by id", errW)
			continue
		}
		response = append(response, dto.GetInventoryResponse{
			UUID:         inventorySnapshot.ID.Hex(),
			BranchID:     inventorySnapshot.BranchID,
			ItemID:       inventorySnapshot.ItemID,
			ItemName:     item.Name,
			ItemCategory: item.Category,
			ItemUnit:     item.Unit,
			CurrentStock: inventorySnapshot.Balance,
			Price:        inventorySnapshot.Latest,
			BranchName:   branch.Name,
		})
	}

	return
}

func (i *inventoryUsecase) mapExistInventoryToRequest(
	requestInventorySnaphots []string,
	inventorySnapshots []model.InventorySnapshot,
) (missingItems []string) {
	mappedItemIdToInventory := make(map[string]model.InventorySnapshot)

	for _, inventorySnapshot := range inventorySnapshots {
		mappedItemIdToInventory[inventorySnapshot.ItemID] = inventorySnapshot
	}

	for _, itemID := range requestInventorySnaphots {
		if _, exist := mappedItemIdToInventory[itemID]; !exist {
			missingItems = append(missingItems, itemID)
		}
	}
	return
}

func (i *inventoryUsecase) BulkCreate(ctx context.Context, items []string, branchID string, endTime time.Time) (errW *error_wrapper.ErrorWrapper) {
	for _, item := range items {
		balance, price, errW := i.inventoryDomain.CalculatePriceAndBalance(ctx, endTime, item, branchID, nil)
		if errW != nil {
			fmt.Println("Error calculating price and balance in bulk create", errW)
			continue
		}

		errW = i.inventorySnapshotDomain.Upsert(ctx, dto.CreateInventorySnapshotRequest{
			ItemID:   item,
			BranchID: branchID,
			Value:    price,
			Balance:  balance,
			Date:     endTime,
		})
		if errW != nil {
			fmt.Println("Error upserting inventory snapshot", errW)
			continue
		}
	}
	return
}

func (i *inventoryUsecase) RecalculateInventory(ctx context.Context, payload dto.RecalculateInventoryRequest) (errW *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.RecalculateInventory(ctx, payload)
}
