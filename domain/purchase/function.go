package purchase

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (p *purchaseDomain) Create(payload dto.CreatePurchaseRequest, userID string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase := model.Purchase{
		SupplierID:   payload.SupplierID,
		BranchID:     payload.BranchID,
		ItemID:       payload.ItemID,
		Quantity:     payload.Quantity,
		PurchaseCost: payload.PurchaseCost,
		Unit:         payload.Unit,
	}

	// 1. Create the purchase record first
	createdPurchase, errW := p.purchaseResource.Create(payload.SupplierID, purchase)
	if errW != nil {
		return nil, errW
	}

	// 2. Create stock transaction
	errW = p.stockTransactionResource.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            userID,
		Quantity:            payload.Quantity,
		Cost:                payload.PurchaseCost,
		Unit:                payload.Unit,
	})
	if errW != nil {
		return nil, errW
	}

	// 3. Handle branch item inventory
	errW = p.syncBranchItemInventory(context.Background(), payload.BranchID, payload.ItemID)
	if errW != nil {
		return nil, errW
	}

	return createdPurchase, nil
}

func (p *purchaseDomain) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindAll()
}

func (p *purchaseDomain) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindByID(id)
}

func (p *purchaseDomain) Update(id, supplierId, branchId, itemId string, quantity float64, purchaseCost float64) *error_wrapper.ErrorWrapper {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}
	return p.purchaseResource.Update(id, purchase)
}

func (p *purchaseDomain) Delete(ctx context.Context, id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	// 1. Delete the purchase and get the deleted data
	deletedPurchase, errW := p.purchaseResource.Delete(id)
	if errW != nil {
		return nil, errW
	}

	errW = p.stockTransactionResource.Create(model.StockTransaction{
		BranchOriginID:      deletedPurchase.BranchID,
		BranchDestinationID: deletedPurchase.BranchID,
		ItemID:              deletedPurchase.ItemID,
		Type:                "OUT",
		IssuerID:            "",
		Quantity:            deletedPurchase.Quantity,
		Cost:                deletedPurchase.PurchaseCost,
		Unit:                deletedPurchase.Unit,
	})

	if errW != nil {
		return nil, errW
	}

	// 2. Sync branch item inventory after deletion
	errW = p.syncBranchItemInventory(ctx, deletedPurchase.BranchID, deletedPurchase.ItemID)
	if errW != nil {
		return nil, errW
	}

	return deletedPurchase, nil
}

func (p *purchaseDomain) syncBranchItemInventory(ctx context.Context, branchID, itemID string) *error_wrapper.ErrorWrapper {
	branchItem, errW := p.branchItemResource.FindByBranchAndItem(branchID, itemID)

	if errW != nil && errW.Is(model.RErrDataNotFound) {
		errW = p.branchItemResource.Create(model.BranchItem{
			BranchID:     branchID,
			ItemID:       itemID,
			CurrentStock: 0,
		})

		if errW != nil {
			return errW
		}
	}
	// Update existing branch item
	currentBalance, errW := p.calculateCurrentBalance(ctx, branchID, itemID)
	if errW != nil {
		return errW
	}

	currentPrice, errW := p.calculatePrice(ctx, branchID, itemID, currentBalance)
	if errW != nil {
		return errW
	}

	_, errW = p.branchItemResource.Update(ctx, model.BranchItem{
		UUID:         branchItem.UUID,
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: currentBalance,
		Price:        currentPrice,
	})

	return errW
}

// calculateCurrentBalance calculates current stock balance from all stock transactions
func (p *purchaseDomain) calculateCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper) {
	allTransactions, err := p.stockTransactionResource.FindAll()
	if err != nil {
		return 0.0, err
	}

	item, errW := p.itemResource.FindByID(itemID)
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

// calculatePrice calculates average price based on recent purchases using FIFO
func (p *purchaseDomain) calculatePrice(ctx context.Context, branchID, itemID string, currentBalance float64) (float64, *error_wrapper.ErrorWrapper) {
	limit := 10
	offset := 0

	purchaseStock := 0.0
	var (
		allPurchases []model.Purchase
		item         model.Item
	)

	for purchaseStock < currentBalance {
		purchases, errW := p.purchaseResource.FindByBranchAndItem(branchID, itemID, offset, limit)
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
