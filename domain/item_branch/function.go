package item_branch

import (
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (s *itemBranchDomain) Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	stockBalance := model.ItemBranch{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: 0.0,
	}

	return s.itemBranchResource.Create(stockBalance)
}

func (s *itemBranchDomain) FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchResource.FindAll()
}

func (s *itemBranchDomain) FindByBranch(branchID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchResource.FindByBranch(branchID)
}

func (s *itemBranchDomain) FindByItem(itemID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchResource.FindByItem(itemID)
}

func (s *itemBranchDomain) FindByBranchAndItem(branchID, itemID string) (*model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchResource.FindByBranchAndItem(branchID, itemID)
}

func (s *itemBranchDomain) Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper {
	return s.itemBranchResource.Update(branchID, itemID, currentStock)
}

func (s *itemBranchDomain) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	return s.itemBranchResource.Delete(branchID, itemID)
}

func (s *itemBranchDomain) SyncCurrentBalance(branchID, itemID string) *error_wrapper.ErrorWrapper {
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

	return s.itemBranchResource.Update(branchID, itemID, totalBalance)
}
