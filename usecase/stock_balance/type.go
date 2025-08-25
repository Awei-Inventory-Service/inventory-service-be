package stockbalance

import (
	stockbalance "github.com/inventory-service/domain/stock_balance"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type StockBalanceUsecase interface {
	FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.StockBalance, *error_wrapper.ErrorWrapper)
}

type stockBalanceUsecase struct {
	stockBalanceDomain stockbalance.StockBalanceDomain
}
