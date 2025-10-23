package inventory

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (i *inventoryDomain) Create(branchID, itemID string, currentStock int) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	branchItem := model.Inventory{
		BranchID: branchID,
		ItemID:   itemID,
	}

	return i.inventoryResource.Create(branchItem)
}

func (i *inventoryDomain) FindAll() (results []dto.GetInventoryResponse, errW *error_wrapper.ErrorWrapper) {
	branchItems, errW := i.inventoryResource.FindAll()

	if errW != nil {
		return
	}
	for _, branchItem := range branchItems {
		results = append(results, dto.GetInventoryResponse{
			UUID:         branchItem.UUID,
			BranchID:     branchItem.BranchID,
			BranchName:   branchItem.Branch.Name,
			ItemID:       branchItem.ItemID,
			ItemName:     branchItem.Item.Name,
			ItemCategory: branchItem.Item.Category,
			CurrentStock: branchItem.Stock,
			Price:        branchItem.Value,
			ItemUnit:     branchItem.Item.Unit,
		})

	}
	return
}

func (i *inventoryDomain) FindByBranch(branchID string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryResource.FindByBranch(branchID)
}

func (i *inventoryDomain) FindByItem(itemID string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryResource.FindByItem(itemID)
}

func (i *inventoryDomain) FindByBranchAndItem(branchID, itemID string) (*model.Inventory, *error_wrapper.ErrorWrapper) {

	return i.inventoryResource.FindByBranchAndItem(branchID, itemID)
}

func (i *inventoryDomain) Update(ctx context.Context, payload model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryResource.Update(ctx, payload)
}

func (i *inventoryDomain) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	return i.inventoryResource.Delete(branchID, itemID)
}

func (i *inventoryDomain) SyncCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper) {
	allTransactions, err := i.stockTransactionResource.FindAll()
	if err != nil {
		return 0.0, err
	}

	item, errW := i.itemResource.FindByID(itemID)
	if errW != nil {
		return 0.0, errW
	}

	var totalBalance float64
	for _, transaction := range allTransactions {
		if transaction.ItemID != itemID {
			continue
		}
		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, item.Unit)

		if transaction.Type == "IN" && transaction.BranchDestinationID == branchID {
			totalBalance += balance
		} else if transaction.Type == "OUT" && transaction.BranchOriginID == branchID {
			totalBalance -= balance
		}
	}

	return totalBalance, nil
}

func (i *inventoryDomain) SyncBranchItem(ctx context.Context, branchID, itemID string) (currentStock, currentPrice float64, errW *error_wrapper.ErrorWrapper) {
	var (
		branchItem *model.Inventory
	)

	branchItem, errW = i.inventoryResource.FindByBranchAndItem(branchID, itemID)

	if errW != nil && errW.Is(model.RErrDataNotFound) {
		errW = nil
		branchItem, errW = i.inventoryResource.Create(model.Inventory{
			BranchID: branchID,
			ItemID:   itemID,
		})
		fmt.Println("Done creating new branch item", branchItem)
		if errW != nil {
			return
		}
	}
	// Update existing branch item
	currentBalance, errW := i.calculateCurrentBalance(ctx, branchID, itemID)
	if errW != nil {
		return
	}
	fmt.Println("Current balance", currentBalance)
	currentPrice, errW = i.calculatePrice(ctx, branchID, itemID, currentBalance)
	if errW != nil {
		return
	}
	fmt.Println("Current price", currentPrice)
	_, errW = i.inventoryResource.Update(ctx, model.Inventory{
		UUID:     branchItem.UUID,
		BranchID: branchID,
		ItemID:   itemID,
		Stock:    currentBalance,
		Value:    currentPrice,
	})

	if errW != nil {
		fmt.Println("Error updating inventory", errW)
		return
	}

	errW = i.inventorySnapshotResource.Upsert(ctx, dto.CreateInventorySnapshotRequest{
		ItemID:   itemID,
		BranchID: branchID,
		Value:    currentPrice,
	})

	if errW != nil {
		fmt.Println("Error upserting inventory snapshot", errW)
		return
	}

	return
}

func (b *inventoryDomain) calculateCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper) {

	allTransactions, err := b.stockTransactionResource.FindWithFilter([]map[string]interface{}{
		{
			"field": "item_id",
			"value": itemID,
		},
		{
			"field": "deleted_at",
			"value": nil,
		},
	}, "created_at DESC", 0, 0)

	if err != nil {
		return 0.0, err
	}

	item, errW := b.itemResource.FindByID(itemID)
	if errW != nil {
		return 0.0, errW
	}

	var totalBalance float64
	for _, transaction := range allTransactions {
		if transaction.ItemID != itemID {
			continue
		}
		fmt.Println("INi transaction", transaction.Unit, item.Unit, transaction.Quantity)
		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, item.Unit)
		fmt.Println("Ini balance dan transaction type", balance, transaction.Type, branchID, transaction.BranchOriginID)
		if transaction.Type == "IN" && transaction.BranchDestinationID == branchID {
			totalBalance += balance
		} else if transaction.Type == "OUT" && transaction.BranchOriginID == branchID {
			totalBalance -= balance
		}
	}
	fmt.Println("INI TOTAL BALANCE", totalBalance)
	return totalBalance, nil
}

// calculatePrice calculates average price based on recent purchases using FIFO
func (b *inventoryDomain) calculatePrice(ctx context.Context, branchID, itemID string, currentBalance float64) (float64, *error_wrapper.ErrorWrapper) {
	limit := 10
	offset := 0

	transactionStock := 0.0
	var (
		allStockTransactions []model.StockTransaction
	)

	for transactionStock < currentBalance {
		stockTransactions, errW := b.stockTransactionResource.FindWithFilter([]map[string]interface{}{
			{
				"field": "item_id",
				"value": itemID,
			},
			{
				"field": "deleted_at",
				"value": nil,
			},
			{
				"field": "type",
				"value": "IN",
			},
		}, "created_at ASC", limit, offset)

		if len(stockTransactions) == 0 {
			break
		}

		if errW != nil {
			return 0.0, errW
		}

		for _, transaction := range stockTransactions {
			allStockTransactions = append(allStockTransactions, transaction)
			transactionQuantity := utils.StandarizeMeasurement(transaction.Quantity, transaction.Unit, transaction.Item.Unit)
			transactionStock += transactionQuantity

			if transactionStock >= currentBalance {
				break
			}
		}

		offset += limit
	}

	if len(allStockTransactions) == 0 {
		return 0.0, nil
	}

	totalPrice := 0.0
	totalItem := 0.0

	for _, transaction := range allStockTransactions {
		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, transaction.Item.Unit)
		totalItem += balance
		totalPrice += transaction.Cost
	}

	// Prevent division by zero which causes NaN
	if totalItem == 0 {
		return 0.0, nil
	}

	avgPrice := totalPrice / totalItem

	return avgPrice, nil
}

func (i *inventoryDomain) BulkSyncBranchItems(ctx context.Context, branchID string, itemIDs []string) *error_wrapper.ErrorWrapper {
	for _, itemId := range itemIDs {
		_, _, errW := i.SyncBranchItem(ctx, branchID, itemId)

		if errW != nil {
			return errW
		}
	}
	return nil
}

func (i *inventoryDomain) GetPrice(ctx context.Context, date dto.CustomDate, itemID, branchID string) (price float64, errW *error_wrapper.ErrorWrapper) {

	inventorySnapshot, errW := i.inventorySnapshotResource.Get(ctx, []dto.Filter{
		{
			Key:    "day",
			Values: []string{fmt.Sprintf("%d", date.Day)},
		},
		{
			Key:    "month",
			Values: []string{fmt.Sprintf("%d", date.Month)},
		},
		{
			Key:    "year",
			Values: []string{fmt.Sprintf("%d", date.Year)},
		},
		{
			Key:    "item_id",
			Values: []string{itemID},
		},
		{
			Key:    "branch_id",
			Values: []string{branchID},
		},
	}, []dto.Order{}, 1, 0)

	if errW != nil {
		return
	}

	if len(inventorySnapshot) == 0 {
		errW = error_wrapper.New(model.RErrDataNotFound, "No inventory snapshot found")
		return
	}

	snapshot := inventorySnapshot[0]
	snapshot.SortValuesBasedOnTimestamp()

	if len(snapshot.Values) == 0 {
		errW = error_wrapper.New(model.RErrDataNotFound, "No price values found in inventory snapshot")
		return
	}

	price = snapshot.Values[0].Value
	return
}
