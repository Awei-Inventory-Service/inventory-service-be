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
	fmt.Printf("Syncing branch: %s and item with id: %s\n", branchID, itemID)
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
		Balance:  currentBalance,
		Date:     time.Now(),
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
			{
				"field": "branch_destination_id",
				"value": branchID,
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

func (i *inventoryDomain) GetInventoryByDate(ctx context.Context, date dto.CustomDate, itemID, branchID string) (response dto.GetInventoryPriceAndValueByDate, errW *error_wrapper.ErrorWrapper) {

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
		_, _, errW = i.SyncBranchItem(ctx, branchID, itemID)
		if errW != nil {
			return
		}
		return
	}

	snapshot := inventorySnapshot[0]

	response.Price = snapshot.Latest
	response.Balance = snapshot.Balance
	response.ItemID = inventorySnapshot[0].ItemID
	return
}

func (i *inventoryDomain) Get(ctx context.Context, payload dto.GetListRequest) (inventories []dto.GetInventoryResponse, errW *error_wrapper.ErrorWrapper) {
	inventoriesRaw, errW := i.inventoryResource.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
	if errW != nil {
		fmt.Println("Error getting all inventories resource", errW)
		return
	}

	for _, inventory := range inventoriesRaw {
		inventories = append(inventories, dto.GetInventoryResponse{
			UUID:         inventory.UUID,
			BranchID:     inventory.Branch.UUID,
			BranchName:   inventory.Branch.Name,
			ItemID:       inventory.Item.UUID,
			ItemName:     inventory.Item.Name,
			ItemCategory: inventory.Item.Category,
			ItemUnit:     inventory.Item.Unit,
			CurrentStock: inventory.Stock,
			Price:        inventory.Value,
		})
	}
	return
}

// Old purcahse date  time.now() -> salah
func (i *inventoryDomain) BuildInventorySnapshotFilter(newPurchaseDate, oldPurchaseDate time.Time) (filters []dto.Filter) {
	if newPurchaseDate.After(oldPurchaseDate) {
		filters = append(filters, []dto.Filter{
			{
				Key:      "date",
				Wildcard: ">=",
				Values:   []string{utils.StartOfDay(oldPurchaseDate).Format("2006-01-02 15:04:05")},
			},
			{
				Key:      "date",
				Wildcard: "<=",
				Values:   []string{utils.EndOfDay(newPurchaseDate).Format("2006-01-02 15:04:05")},
			},
		}...)
	} else {
		filters = append(filters, []dto.Filter{
			{
				Key:      "date",
				Wildcard: ">=",
				Values:   []string{utils.EndOfDay(oldPurchaseDate).Format("2006-01-02 15:04:05")},
			},
			{
				Key:      "date",
				Wildcard: "<=",
				Values:   []string{utils.StartOfDay(newPurchaseDate).Format("2006-01-02 15:04:05")},
			},
		}...)
	}
	return
}

func (i *inventoryDomain) RecalculateInventory(ctx context.Context, payload dto.RecalculateInventoryRequest) (errW *error_wrapper.ErrorWrapper) {
	// Parse the startTime and check if it's today
	parsedTime, err := time.Parse("2006-01-02", payload.NewTime)
	if err != nil {
		return error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
	}

	today := time.Now().Truncate(24 * time.Hour)
	if parsedTime.Equal(today) {
		fmt.Println("Skipping recalculate inventory because start time equal to today, no need to recalculte anything")
		return nil
	}

	var (
		endTime = time.Now()
	)

	startTime := parsedTime

	if payload.PreviousTime != nil && payload.PreviousTime.Before(parsedTime) {
		startTime = *payload.PreviousTime
	}

	fmt.Println("Ini start time", startTime)
	filter := []dto.Filter{
		{
			Key:    "branch_id",
			Values: []string{payload.BranchID},
		},
		{
			Key:    "item_id",
			Values: []string{payload.ItemID},
		},
		{
			Key:      "date",
			Wildcard: ">=",
			Values:   []string{utils.StartOfDay(startTime).Format("2006-01-02 15:04:05")},
		},
		{
			Key:      "date",
			Wildcard: "<=",
			Values:   []string{utils.EndOfDay(endTime).Format("2006-01-02 15:04:05")},
		},
	}

	order := []dto.Order{
		{
			Key:   "date",
			IsAsc: true,
		},
	}

	previousSnapshot, errW := i.BackFillInventorySnapshot(ctx, startTime, payload.BranchID, payload.ItemID)
	if errW != nil {
		fmt.Println("Error getting previous day snapshot", errW)
		return
	}

	// Check if inventory snapshot exist for current date
	_, errW = i.inventorySnapshotResource.GetSnapshotBasedOndDate(ctx, parsedTime)
	if errW != nil && errW.Is(model.RErrDataNotFound) {
		// Create new empty one
		errW = i.inventorySnapshotResource.Upsert(ctx, dto.CreateInventorySnapshotRequest{
			BranchID: payload.BranchID,
			ItemID:   payload.ItemID,
			Value:    0,
			Balance:  0,
			Date:     parsedTime,
		})
		if errW != nil {
			fmt.Println("Error creating new intentory resource", errW)
		}
	}

	fmt.Println("Ini filter di recalculate inventory", filter)
	// Get all inventory snapshots need to be updated
	inventorySnasphots, errW := i.inventorySnapshotResource.Get(ctx, filter, order, 0, 0)
	if errW != nil {
		fmt.Println("Error getting inventory snapshout", errW)
		return
	}
	fmt.Println("Ini inventory snapshots", len(inventorySnasphots), inventorySnasphots)

	for _, inventorySnapshot := range inventorySnasphots {
		fmt.Println("Ini inventory snapshot", inventorySnapshot)
		fmt.Println("Ini previous snapshot", previousSnapshot)

		startDate := utils.StartOfDay(inventorySnapshot.Date)
		endDate := utils.EndOfDay(inventorySnapshot.Date)
		// Update balance nya
		balance, price, errW := i.CalculatePriceAndBalance(ctx, startDate, payload.ItemID, payload.BranchID, &endDate)
		if errW != nil {
			fmt.Println("Error calculating price and balance", errW)
			return errW
		}
		fmt.Println("iNi balance dan price", balance, price)
		totalBalance := previousSnapshot.Balance + balance

		newPrice := 0.0

		if totalBalance > 0 {
			newPrice = (price*balance + previousSnapshot.Balance*previousSnapshot.Latest) / totalBalance
		}

		inventorySnapshot.Balance = totalBalance
		inventorySnapshot.Latest = newPrice
		inventorySnapshot.Values = append(inventorySnapshot.Values, struct {
			Timestamp time.Time "json:\"timestamp\""
			Value     float64   "json:\"value\""
		}{
			Timestamp: inventorySnapshot.ID.Timestamp(),
			Value:     newPrice,
		})
		fmt.Println("Ini updated inventory", inventorySnapshot)
		errW = i.inventorySnapshotResource.Update(ctx, inventorySnapshot.ID.Hex(), inventorySnapshot)
		if errW != nil {
			fmt.Println("Error updating inventory resource", errW)
			return errW
		}
		previousSnapshot = inventorySnapshot
	}
	return
}

func (i *inventoryDomain) CalculatePriceAndBalance(ctx context.Context, endTime time.Time, itemID, branchID string, startTime *time.Time) (balance, price float64, errW *error_wrapper.ErrorWrapper) {
	// 1. Get stock transactions until the end time
	endOfDay := utils.EndOfDay(endTime)

	filter := []dto.Filter{
		{
			Key:    "item_id",
			Values: []string{itemID},
		},
		{
			Key:      "transaction_date",
			Wildcard: "<=",
			Values:   []string{endOfDay.Format("2006-01-02 15:04:05")},
		},
		{
			Key:    "deleted_at",
			Values: []string{""},
		},
	}

	// Add startTime filter only if provided
	if startTime != nil {
		filter = append(filter, dto.Filter{
			Key:      "transaction_date",
			Wildcard: ">=",
			Values:   []string{startTime.Format("2006-01-02 15:04:05")},
		})
	}

	fmt.Println("Ini filter", filter)
	stockTransactions, errW := i.stockTransactionResource.Get(ctx, filter, []dto.Order{}, 0, 0)
	if errW != nil {
		fmt.Println("Error finding stock transaction with filter", errW)
		return
	}

	var (
		totalBalance, totalPrice, avg float64
	)

	for _, stockTransaction := range stockTransactions {
		balance := utils.StandarizeMeasurement(stockTransaction.Quantity, stockTransaction.Unit, stockTransaction.Item.Unit)

		if stockTransaction.BranchDestinationID != branchID && stockTransaction.BranchOriginID != branchID {
			continue
		}
		fmt.Println("Ini stock transaction", stockTransaction)
		if stockTransaction.Type == "IN" {
			totalBalance += balance
			totalPrice += stockTransaction.Cost
		} else if stockTransaction.Type == "OUT" {
			totalBalance -= balance
		}
	}

	if len(stockTransactions) > 0 {
		avg = totalPrice / totalBalance
	}

	return totalBalance, avg, nil
}

func (i *inventoryDomain) BackFillInventorySnapshot(ctx context.Context, startTime time.Time, branchID, itemID string) (inventorySnapshot model.InventorySnapshot, errW *error_wrapper.ErrorWrapper) {

	// Get 1 snapshot before the date
	previousSnapshot, errW := i.inventorySnapshotResource.GetPreviousDaySnapshot(ctx, startTime, branchID, itemID)
	if errW != nil {
		if errW.Is(model.RErrDataNotFound) {
			previousDay := startTime.AddDate(0, 0, -1)

			balance, price, errW := i.CalculatePriceAndBalance(ctx, previousDay, itemID, branchID, nil)
			if errW != nil {
				fmt.Println("Error calculating price and balance for previous day", errW)
				return inventorySnapshot, errW
			}

			newInventorySnaspshot := model.InventorySnapshot{
				ItemID:   itemID,
				BranchID: branchID,
				Average:  price,
				Latest:   price,
				Balance:  balance,
				Date:     previousDay,
				Day:      previousDay.Day(),
				Month:    int(previousDay.Month()),
				Year:     previousDay.Year(),
				Values: []struct {
					Timestamp time.Time "json:\"timestamp\""
					Value     float64   "json:\"value\""
				}{{
					Timestamp: previousDay,
					Value:     price,
				}},
			}
			errW = i.inventorySnapshotResource.Create(ctx, newInventorySnaspshot)

			if errW != nil {
				fmt.Println("Error creating new inventory snapshot", errW)
				return model.InventorySnapshot{}, errW
			}
			return newInventorySnaspshot, nil
		}
		return model.InventorySnapshot{}, errW
	}

	return *previousSnapshot, errW
}
