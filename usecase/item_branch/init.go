package item_branch

import itembranch "github.com/inventory-service/domain/item_branch"

func NewStockBalanceUsecase(itemBranchDomain itembranch.ItemBranchDomain) ItemBranchUsecase {
	return &itemBranchUsecase{
		itemBranchDomain: itemBranchDomain,
	}
}
