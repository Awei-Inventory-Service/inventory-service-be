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
		if !errW.Is(model.RErrDataNotFound) {
			return 0, nil, errW
		}
		nextPurchaseChain, errW := i.itemPurchaseChainRepository.Get(ctx, model.ItemPurchaseChain{
			ItemID:   itemID,
			BranchID: branchID,
			Status:   model.StatusNotUsed,
		})

		if errW != nil {
			return 0, nil, errW
		}
		nextPurchaseChain[0].Status = model.StatusInUse

		errW = i.itemPurchaseChainRepository.Update(ctx, nextPurchaseChain[0].ID, model.ItemPurchaseChain{
			ItemID:   nextPurchaseChain[0].ItemID,
			BranchID: nextPurchaseChain[0].BranchID,
			Status:   model.StatusInUse,
			Purchase: nextPurchaseChain[0].Purchase,
			Sales:    nextPurchaseChain[0].Sales,
		})

		if errW != nil {
			fmt.Println("Error di update")
			return 0, nil, errW
		}
		purchaseChain = nextPurchaseChain
	}

	quantityLeft := quantity
	for quantityLeft > 0 {
		if purchaseChain[0].Quantity < quantityLeft {
			cost += float64(purchaseChain[0].Quantity) * purchaseChain[0].Purchase.Item.Price
			quantityLeft -= purchaseChain[0].Quantity
			purchaseChain[0].Quantity = 0
			purchaseChain[0].Status = model.StatusUsed
			results = append(results, purchaseChain[0])
			nextPurchaseChain, errW := i.itemPurchaseChainRepository.Get(
				ctx,
				model.ItemPurchaseChain{
					ItemID:   itemID,
					BranchID: branchID,
					Status:   model.StatusNotUsed,
				},
			)

			if errW != nil {
				return 0, nil, errW
			}
			nextPurchaseChain[0].Status = model.StatusInUse
			purchaseChain[0] = nextPurchaseChain[0]
		} else {
			cost += float64(quantityLeft) * purchaseChain[0].Purchase.Item.Price
			purchaseChain[0].Quantity -= quantityLeft

			if purchaseChain[0].Quantity == 0 {
				purchaseChain[0].Status = model.StatusUsed
			}
			quantityLeft = 0
			results = append(results, purchaseChain[0])
		}
	}
	return cost, results, nil
}
