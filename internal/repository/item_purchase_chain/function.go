package itempurchasechain

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemPurchaseChainRepository) Create(ctx context.Context, itemID string, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper {
	itemPurchaseChain := model.ItemPurchaseChain{
		ItemID:       itemID,
		BranchID:     branchID,
		Purchase:     purchase,
		Quantity:     purchase.Quantity,
		Status:       model.StatusNotUsed,
		SalesRecords: []model.Sales{},
	}

	_, err := i.itemPurchaseChainCollection.InsertOne(
		ctx,
		itemPurchaseChain,
	)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return nil
}

// Find By Item Id dan BranchId

// Update -> bisa update quantity / status

// CalculateCost -> terima parameter item id sama berapa banyak butuhnya. Return []item_purchase_chain. Dipake sama sales nanti buat ngisi cost

// Logic : 
// 1. Cari berdasarkan item id dan branch id + status = in use 
// 2. IF no 1 cukup -> cost = quantity di request * price purchase. Update stock dan status nomor 1
// 3. If no 1 ga cukup -> cost = quantity no 1 * price no 1 + quantity no 2 * price no 2. Update stock dan status di nomor 1 dan nomor 2
