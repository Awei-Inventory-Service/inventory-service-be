package stockbalance

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockBalanceDomain) Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	stockBalance := model.StockBalance{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: currentStock,
	}

	return s.stockBalanceResource.Create(stockBalance)
}

func (s *stockBalanceDomain) FindAll() ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceResource.FindAll()
}

func (s *stockBalanceDomain) FindByBranch(branchID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceResource.FindByBranch(branchID)
}

func (s *stockBalanceDomain) FindByItem(itemID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceResource.FindByItem(itemID)
}

func (s *stockBalanceDomain) FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, *error_wrapper.ErrorWrapper) {
	return s.stockBalanceResource.FindByBranchAndItem(branchID, itemID)
}

func (s *stockBalanceDomain) Update(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	return s.stockBalanceResource.Update(branchID, itemID, currentStock)
}

func (s *stockBalanceDomain) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	return s.stockBalanceResource.Delete(branchID, itemID)
}

func (s *stockBalanceDomain) SyncCurrentBalance(branchID, itemID string) *error_wrapper.ErrorWrapper {
	allTransactions, err := s.stockTransactionResource.FindAll()
	if err != nil {
		return err
	}

	var totalBalance int
	for _, transaction := range allTransactions {
		if transaction.ItemID != itemID {
			continue
		}

		if transaction.Type == "IN" && transaction.BranchDestinationID == branchID {
			totalBalance += transaction.Quantity
		} else if transaction.Type == "OUT" && transaction.BranchOriginID == branchID {
			totalBalance -= transaction.Quantity
		}
	}

	return s.stockBalanceResource.Update(branchID, itemID, totalBalance)
}
