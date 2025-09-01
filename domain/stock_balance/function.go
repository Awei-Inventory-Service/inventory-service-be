package stockbalance

import (
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (s *stockBalanceDomain) Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	stockBalance := model.StockBalance{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: 0.0,
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

func (s *stockBalanceDomain) Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper {
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

	item, errW := s.itemResource.FindByID(itemID)

	if errW != nil {
		return errW
	}

	var totalBalance float64
	for _, transaction := range allTransactions {
		if transaction.ItemID != itemID {
			continue
		}
		fmt.Println("INI TRANSACTION QUANTITY UNIT", transaction.Quantity, transaction.Unit, item.Unit)
		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, item.Unit)
		fmt.Println("INI VALANCE", balance)
		if transaction.Type == "IN" && transaction.BranchDestinationID == branchID {
			totalBalance += balance
		} else if transaction.Type == "OUT" && transaction.BranchOriginID == branchID {
			totalBalance -= balance
		}
	}

	return s.stockBalanceResource.Update(branchID, itemID, totalBalance)
}
