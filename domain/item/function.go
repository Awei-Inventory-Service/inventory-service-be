package item

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/utils"
	"github.com/inventory-service/model"
)

func (i *itemDomain) Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.Create(item)
}

func (i *itemDomain) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindAll()
}

func (i *itemDomain) FindByID(ctx context.Context, id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	// itemPrice, errW := i.calculatePrice(ctx, id)
	// fmt.Println("INI ITEMPRICE", itemPrice)
	// if errW != nil {
	// 	return nil, errW
	// }

	item, errW := i.itemResource.FindByID(id)

	if errW != nil {
		return nil, errW
	}

	return item, nil
}

func (i *itemDomain) Update(ctx context.Context, payload dto.UpdateItemRequest, itemID string) *error_wrapper.ErrorWrapper {
	itemCategory, errW := utils.ParseItemCategory(payload.Category)

	if errW != nil {
		return errW
	}

	item := model.Item{
		UUID:       itemID,
		Name:       payload.Name,
		Category:   itemCategory,
		Price:      payload.Price,
		Unit:       payload.Unit,
		SupplierID: &payload.SupplierID,
	}

	updatedItem, errW := i.itemResource.Update(ctx, item)

	if errW != nil {
		return errW
	}

	errW = i.itemCompositionResource.DeleteByItemID(ctx, updatedItem.UUID)

	if errW != nil {
		return errW
	}

	for _, itemComposition := range payload.ItemCompositions {
		errW = i.itemCompositionResource.Create(ctx, model.ItemComposition{
			ParentItemID: updatedItem.UUID,
			ChildItemID:  itemComposition.ItemID,
			Ratio:        itemComposition.Ratio,
			Notes:        itemComposition.Notes,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if errW != nil {
			fmt.Println("Error creating item composition", errW.ActualError())
			return errW
		}
	}
	return nil
}

func (i *itemDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return i.itemResource.Delete(id)
}

func (i *itemDomain) calculatePrice(ctx context.Context, itemID string) (float64, *error_wrapper.ErrorWrapper) {
	branchId := ctx.Value("branch_id")
	limit := 10
	offset := 0

	stockBalance, errW := i.stockBalanceResource.FindByBranchAndItem(fmt.Sprint(branchId), itemID)
	if errW != nil {
		return 0.0, errW
	}

	purchaseStock := 0.0
	var allPurchases []model.Purchase

	for purchaseStock < stockBalance.CurrentStock {
		purchases, errW := i.purchaseResource.FindByBranchAndItem(fmt.Sprint(branchId), itemID, offset, limit)
		if errW != nil {
			return 0.0, errW
		}

		if len(purchases) == 0 {
			break
		}

		for _, purchase := range purchases {
			allPurchases = append(allPurchases, purchase)
			purchaseStock += purchase.Quantity
			if purchaseStock >= stockBalance.CurrentStock {
				break
			}
		}

		offset += limit
	}

	totalPrice := 0.0
	totalItem := 0.0

	for _, purchase := range allPurchases {
		totalItem += float64(purchase.Quantity)
		totalPrice += purchase.PurchaseCost
	}

	avgPrice := totalPrice / totalItem

	return avgPrice, nil
}
