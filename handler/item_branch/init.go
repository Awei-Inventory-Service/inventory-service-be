package stockbalance

import itembranch "github.com/inventory-service/usecase/item_branch"

func NewStockBalanceHandler(itemBranchUsecase itembranch.ItemBranchUsecase) ItemBranchHandler {
	return &itemBranchHandler{
		itemBranchUsecase: itemBranchUsecase,
	}
}
