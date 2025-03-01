package itempurchasechain

import (
	"context"
	"fmt"

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

func (i *itemPurchaseChainService) CalculateCost(ctx context.Context, itemID string, branchID string, quantity int) (float64, []model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	var (
		results []model.ItemPurchaseChainGet
		cost    = 0.0
	)
	purchaseChain, errW := i.itemPurchaseChainRepository.Get(ctx, model.ItemPurchaseChain{
		ItemID:   itemID,
		BranchID: branchID,
		Status:   model.StatusInUse,
	})

	if errW != nil {
		return 0, nil, errW
	}
	if purchaseChain[0].Quantity < quantity {
		quantityLeft := quantity - purchaseChain[0].Quantity

		cost += float64(purchaseChain[0].Quantity) * purchaseChain[0].Purchase.Item.Price
		purchaseChain[0].Quantity = 0
		purchaseChain[0].Status = model.StatusUsed
		results = append(results, purchaseChain[0])
		nextPurchaseChain, errW := i.itemPurchaseChainRepository.Get(ctx, model.ItemPurchaseChain{
			ItemID:   itemID,
			BranchID: branchID,
			Status:   model.StatusNotUsed,
		})

		if errW != nil {
			return 0, nil, errW
		}
		fmt.Println("INI COST", cost)
		// TO DO : Edge case kalau 2 item purchase chain masih ga cukup
		fmt.Println("INI QUANTITY LEFT", quantityLeft)
		if nextPurchaseChain[0].Quantity >= quantityLeft {
			nextPurchaseChain[0].Quantity -= quantityLeft
			nextPurchaseChain[0].Status = model.StatusInUse
			cost += (float64(quantityLeft) * nextPurchaseChain[0].Purchase.Item.Price)
			results = append(results, nextPurchaseChain[0])
		}

	} else {
		purchaseChain[0].Quantity = purchaseChain[0].Quantity - quantity
		cost += float64(quantity) * purchaseChain[0].Purchase.Item.Price
		results = append(results, purchaseChain[0])
	}
	return cost, results, nil
}
