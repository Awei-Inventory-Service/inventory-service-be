package stockbalance

import (
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockBalanceUsecase) FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceDomain.FindByBranchAndItem(payload.BranchId, payload.ItemId)
}

func (s *stockBalanceUsecase) FindByBranchId(branchId string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceDomain.FindByBranch(branchId)
}

func (s *stockBalanceUsecase) FindAll() ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceDomain.FindAll()
}
