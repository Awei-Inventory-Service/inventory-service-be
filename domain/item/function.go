package item

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/utils"
	"github.com/inventory-service/model"
)

func (i *itemDomain) Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.Create(item)
}

func (i *itemDomain) FindAll() ([]dto.GetItemsResponse, *error_wrapper.ErrorWrapper) {
	rawItems, errW := i.itemResource.FindAll()

	if errW != nil {
		return nil, errW
	}

	var items []dto.GetItemsResponse

	for _, it := range rawItems {
		items = append(items, *i.mapItemModelToDto(it))
	}
	return items, nil
}

func (i *itemDomain) FindByID(ctx context.Context, id string) (*dto.GetItemsResponse, *error_wrapper.ErrorWrapper) {

	item, errW := i.itemResource.FindByID(id)

	if errW != nil {
		return nil, errW
	}

	return i.mapItemModelToDto(*item), nil
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
		SupplierID: &payload.SupplierID,
	}

	_, errW = i.itemResource.Update(ctx, item)

	if errW != nil {
		return errW
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

	stockBalance, errW := i.inventoryResource.FindByBranchAndItem(fmt.Sprint(branchId), itemID)
	if errW != nil {
		return 0.0, errW
	}

	purchaseStock := 0.0
	var allPurchases []model.Purchase

	for purchaseStock < stockBalance.Stock {
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
			if purchaseStock >= stockBalance.Stock {
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

func (i *itemDomain) mapItemModelToDto(item model.Item) *dto.GetItemsResponse {
	var itemCompositions []dto.GetItemCompositionResponse

	if len(item.Compositions.Compositions) > 0 {
		for _, comp := range item.Compositions.Compositions {
			itemCompositions = append(itemCompositions, dto.GetItemCompositionResponse{
				ChildItemID:   comp.ItemID,
				Unit:          comp.Unit,
				ChildItemName: comp.ItemName,
			})
		}
	}

	return &dto.GetItemsResponse{
		UUID:              item.UUID,
		Name:              item.Name,
		Category:          item.Category,
		Unit:              item.Unit,
		ChildCompositions: itemCompositions,
	}
}
