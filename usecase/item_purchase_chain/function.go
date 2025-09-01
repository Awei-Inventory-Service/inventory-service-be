package itempurchasechain

import (
	"context"
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemPurchaseChainService) Create(ctx context.Context, itemID string, branchID string, purchaseID string) *error_wrapper.ErrorWrapper {
	purchase, errW := i.purchaseDomain.FindByID(purchaseID)

	if errW != nil {
		return errW
	}

	_, errW = i.itemDomain.FindByID(ctx, itemID)

	if errW != nil {
		return errW
	}

	_, errW = i.branchDomain.FindByID(branchID)

	if errW != nil {
		return errW
	}

	return i.itemPurchaseChainDomain.Create(ctx, itemID, branchID, *purchase)
}

func (i *itemPurchaseChainService) CalculateCost(ctx context.Context, itemID string, branchID string, quantity float64) (float64, []model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	var (
		results []model.ItemPurchaseChainGet
		cost    = 0.0
	)
	purchaseChain, errW := i.itemPurchaseChainDomain.Get(ctx, model.ItemPurchaseChain{
		ItemID:   itemID,
		BranchID: branchID,
		Status:   model.StatusInUse,
	})

	if errW != nil {
		if !errW.Is(model.RErrDataNotFound) {
			return 0, nil, errW
		}
		nextPurchaseChain, errW := i.itemPurchaseChainDomain.Get(ctx, model.ItemPurchaseChain{
			ItemID:   itemID,
			BranchID: branchID,
			Status:   model.StatusNotUsed,
		})

		if errW != nil {
			return 0, nil, errW
		}
		nextPurchaseChain[0].Status = model.StatusInUse

		errW = i.itemPurchaseChainDomain.Update(ctx, nextPurchaseChain[0].ID, model.ItemPurchaseChain{
			ItemID:   nextPurchaseChain[0].ItemID,
			BranchID: nextPurchaseChain[0].BranchID,
			Status:   model.StatusInUse,
			Purchase: model.ItemPurchaseChainPurchase{
				UUID:         nextPurchaseChain[0].Purchase.UUID,
				Quantity:     nextPurchaseChain[0].Purchase.Quantity,
				BranchId:     nextPurchaseChain[0].Purchase.BranchId,
				PurchaseCost: nextPurchaseChain[0].Purchase.PurchaseCost,
				ItemId:       nextPurchaseChain[0].Purchase.ItemId,
			},
			Sales: nextPurchaseChain[0].Sales,
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
			cost += float64(purchaseChain[0].Quantity) * purchaseChain[0].Purchase.PurchaseCost
			quantityLeft -= purchaseChain[0].Quantity
			purchaseChain[0].Quantity = 0
			purchaseChain[0].Status = model.StatusUsed
			results = append(results, purchaseChain[0])
			nextPurchaseChain, errW := i.itemPurchaseChainDomain.Get(
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
			cost += float64(quantityLeft) * purchaseChain[0].Purchase.PurchaseCost
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
