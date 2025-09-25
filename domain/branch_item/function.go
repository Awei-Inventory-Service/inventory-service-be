package branch_item

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (s *branchItemDomain) Create(branchID, itemID string, currentStock int) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	branchItem := model.BranchItem{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: 0.0,
	}

	return s.branchItemResource.Create(branchItem)
}

func (s *branchItemDomain) FindAll() (results []dto.GetBranchItemResponse, errW *error_wrapper.ErrorWrapper) {
	branchItems, errW := s.branchItemResource.FindAll()

	if errW != nil {
		return
	}
	for _, branchItem := range branchItems {
		results = append(results, dto.GetBranchItemResponse{
			UUID:         branchItem.UUID,
			BranchID:     branchItem.BranchID,
			BranchName:   branchItem.Branch.Name,
			ItemID:       branchItem.ItemID,
			ItemName:     branchItem.Item.Name,
			ItemCategory: branchItem.Item.Category,
			CurrentStock: branchItem.CurrentStock,
			Price:        branchItem.Price,
			ItemUnit:     branchItem.Item.Unit,
		})

	}
	return
}

func (s *branchItemDomain) FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.FindByBranch(branchID)
}

func (s *branchItemDomain) FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.FindByItem(itemID)
}

func (s *branchItemDomain) FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper) {

	return s.branchItemResource.FindByBranchAndItem(branchID, itemID)
}

func (s *branchItemDomain) Update(ctx context.Context, payload model.BranchItem) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.Update(ctx, payload)
}

func (s *branchItemDomain) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	return s.branchItemResource.Delete(branchID, itemID)
}

func (s *branchItemDomain) SyncCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper) {
	allTransactions, err := s.stockTransactionResource.FindAll()
	if err != nil {
		return 0.0, err
	}

	item, errW := s.itemResource.FindByID(itemID)
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

func (s *branchItemDomain) CalculatePrice(ctx context.Context, branchID, itemID string, currentBalance float64) (float64, *error_wrapper.ErrorWrapper) {
	limit := 10
	offset := 0

	purchaseStock := 0.0
	var (
		allPurchases []model.Purchase
		item         model.Item
	)

	for purchaseStock < currentBalance {
		purchases, errW := s.purchaseResource.FindByBranchAndItem(branchID, itemID, offset, limit)
		if errW != nil {
			return 0.0, errW
		}

		if len(purchases) == 0 {
			break
		}

		for _, purchase := range purchases {
			allPurchases = append(allPurchases, purchase)
			purchaseStock += purchase.Quantity
			if purchaseStock >= currentBalance {
				break
			}
		}

		offset += limit
	}

	if len(allPurchases) == 0 {
		return 0.0, nil
	}

	totalPrice := 0.0
	totalItem := 0.0

	for _, purchase := range allPurchases {
		balance := utils.StandarizeMeasurement(float64(purchase.Quantity), purchase.Unit, purchase.Item.Unit)
		totalItem += balance
		totalPrice += purchase.PurchaseCost
		item = purchase.Item
	}

	// Prevent division by zero which causes NaN
	if totalItem == 0 {
		return 0.0, nil
	}

	avgPrice := totalPrice / totalItem * item.PortionSize

	return avgPrice, nil
}

func (s *branchItemDomain) SyncBranchItem(ctx context.Context, branchID, itemID string) *error_wrapper.ErrorWrapper {
	var (
		branchItem *model.BranchItem
	)
	fmt.Println("BRANCH ID AND ITEM ID", branchID, itemID)
	branchItem, errW := s.branchItemResource.FindByBranchAndItem(branchID, itemID)

	if errW != nil && errW.Is(model.RErrDataNotFound) {
		errW = nil
		branchItem, errW = s.branchItemResource.Create(model.BranchItem{
			BranchID:     branchID,
			ItemID:       itemID,
			CurrentStock: 0,
		})
		fmt.Println("Done creating new branch item", branchItem)
		if errW != nil {
			return errW
		}
	}
	// Update existing branch item
	currentBalance, errW := s.calculateCurrentBalance(ctx, branchID, itemID)
	if errW != nil {
		return errW
	}
	fmt.Println("Current balance", currentBalance)
	currentPrice, errW := s.calculatePrice(ctx, branchID, itemID, currentBalance)
	if errW != nil {
		return errW
	}
	fmt.Println("Current price", currentPrice)
	_, errW = s.branchItemResource.Update(ctx, model.BranchItem{
		UUID:         branchItem.UUID,
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: currentBalance,
		Price:        currentPrice,
	})

	return errW
}

func (b *branchItemDomain) calculateCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper) {
	allTransactions, err := b.stockTransactionResource.FindAll()
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
		fmt.Println("INI transaction", transaction)
		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, item.Unit)

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
func (b *branchItemDomain) calculatePrice(ctx context.Context, branchID, itemID string, currentBalance float64) (float64, *error_wrapper.ErrorWrapper) {
	limit := 10
	offset := 0

	purchaseStock := 0.0
	var (
		allPurchases []model.Purchase
		item         model.Item
	)

	for purchaseStock < currentBalance {
		purchases, errW := b.purchaseResource.FindByBranchAndItem(branchID, itemID, offset, limit)
		if errW != nil {
			return 0.0, errW
		}

		if len(purchases) == 0 {
			break
		}

		for _, purchase := range purchases {
			allPurchases = append(allPurchases, purchase)
			purchaseStock += purchase.Quantity
			if purchaseStock >= currentBalance {
				break
			}
		}

		offset += limit
	}

	if len(allPurchases) == 0 {
		return 0.0, nil
	}

	totalPrice := 0.0
	totalItem := 0.0

	for _, purchase := range allPurchases {
		balance := utils.StandarizeMeasurement(float64(purchase.Quantity), purchase.Unit, purchase.Item.Unit)
		totalItem += balance
		totalPrice += purchase.PurchaseCost
		item = purchase.Item
	}

	// Prevent division by zero which causes NaN
	if totalItem == 0 {
		return 0.0, nil
	}

	avgPrice := totalPrice / totalItem * item.PortionSize

	return avgPrice, nil
}
