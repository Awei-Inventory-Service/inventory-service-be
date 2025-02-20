package itempurchasechain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemPurchaseChainService) Create(ctx context.Context, itemID string, branchID string, purchaseID string) *error_wrapper.ErrorWrapper {
	purchase, errW := i.purchaseRepository.FindByID(purchaseID)

	if errW != nil {
		return errW
	}

	_, errW = i.itemRepository.FindByID(itemID)

	if errW != nil{
		return errW
	}

	_, errW = i.branchRepository.FindByID(branchID)

	if errW != nil{
		return errW
	}

	return i.itemPurchaseChainRepository.Create(ctx, itemID, branchID, *purchase)
}
