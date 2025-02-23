package itempurchasechain

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemPurchaseChainService) Create(ctx context.Context, itemID string, branchID string, purchaseID string) *error_wrapper.ErrorWrapper {
	purchase, errW := i.purchaseRepository.FindByID(purchaseID)

	if errW != nil {
		return errW
	}

	_, errW = i.itemRepository.FindByID(itemID)

	if errW != nil {
		return errW
	}

	_, errW = i.branchRepository.FindByID(branchID)

	if errW != nil {
		return errW
	}

	return i.itemPurchaseChainRepository.Create(ctx, itemID, branchID, *purchase)
}

func (i *itemPurchaseChainService) CalculateCost(ctx context.Context, itemID string, branchID string, quantity int) (int, []model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	var (
		results []model.ItemPurchaseChainGet
		cost    = 0
	)
	purchaseChain, errW := i.itemPurchaseChainRepository.Get(ctx, model.ItemPurchaseChain{
		ItemID:   itemID,
		BranchID: branchID,
		Status:   model.StatusInUse,
	})

	if errW != nil {
		return 0, nil, errW
	}
	results = append(results, purchaseChain[0])
	if purchaseChain[0].Quantity < quantity {
		purchaseChain[0].Quantity = 0
		purchaseChain[0].Status = model.StatusUsed
		quantityLeft := quantity - purchaseChain[0].Quantity
		errW = i.itemPurchaseChainRepository.Update(ctx, purchaseChain[0].ID, model.ItemPurchaseChain{
			ItemID:       purchaseChain[0].ItemID,
			BranchID:     purchaseChain[0].BranchID,
			Purchase:     purchaseChain[0].Purchase,
			Quantity:     0,
			Status:       model.StatusUsed,
			SalesRecords: purchaseChain[0].SalesRecords,
		})
		cost += (purchaseChain[0].Quantity * int(purchaseChain[0].Purchase.Item.Price))
		nextPurchaseChain, errW := i.itemPurchaseChainRepository.Get(ctx, model.ItemPurchaseChain{
			ItemID:   itemID,
			BranchID: branchID,
			Status:   model.StatusNotUsed,
		})

		if errW != nil {
			return 0, nil, errW
		}

		// TO DO : Edge case kalau 2 item purchase chain masih ga cukup

		if nextPurchaseChain[0].Quantity >= quantityLeft {
			errW = i.itemPurchaseChainRepository.Update(ctx, nextPurchaseChain[0].ID, model.ItemPurchaseChain{
				ItemID:       nextPurchaseChain[0].ItemID,
				BranchID:     nextPurchaseChain[0].BranchID,
				Purchase:     nextPurchaseChain[0].Purchase,
				Quantity:     nextPurchaseChain[0].Quantity - quantityLeft,
				Status:       model.StatusInUse,
				SalesRecords: nextPurchaseChain[0].SalesRecords,
			})
			results = append(results, nextPurchaseChain[0])
			cost += quantityLeft * int(nextPurchaseChain[0].Purchase.Item.Price)
		}
	} else {
		purchaseChain[0].Quantity = purchaseChain[0].Quantity - quantity
		cost += (quantity * int(purchaseChain[0].Purchase.Item.Price))
	}
	return cost, results, nil
}
