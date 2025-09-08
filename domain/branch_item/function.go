package branch_item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (s *branchItemDomain) Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	branchItem := model.BranchItem{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: 0.0,
	}

	return s.branchItemResource.Create(branchItem)
}

func (s *branchItemDomain) FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.FindAll()
}

func (s *branchItemDomain) FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.FindByBranch(branchID)
}

func (s *branchItemDomain) FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemResource.FindByItem(itemID)
}

func (s *branchItemDomain) FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	
	return s.branchItemResource.FindByBranchAndItem(branchID, itemID)
}

func (s *branchItemDomain) Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper {
	return s.branchItemResource.Update(branchID, itemID, currentStock)
}

func (s *branchItemDomain) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	return s.branchItemResource.Delete(branchID, itemID)
}

func (s *branchItemDomain) SyncCurrentBalance(branchID, itemID string) *error_wrapper.ErrorWrapper {
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

		balance := utils.StandarizeMeasurement(float64(transaction.Quantity), transaction.Unit, item.Unit)

		if transaction.Type == "IN" && transaction.BranchDestinationID == branchID {
			totalBalance += balance
		} else if transaction.Type == "OUT" && transaction.BranchOriginID == branchID {
			totalBalance -= balance
		}
	}

	return s.branchItemResource.Update(branchID, itemID, totalBalance)
}
