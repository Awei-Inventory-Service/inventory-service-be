package stocktransaction

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockTransactionDomain) Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Create(transaction)
}

func (s *stockTransactionDomain) FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindAll()
}

func (s *stockTransactionDomain) FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindByID(id)
}

func (s *stockTransactionDomain) Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Update(id, transaction)
}

func (s *stockTransactionDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Delete(id)
}

func (s *stockTransactionDomain) FindWithFilter(filters []map[string]interface{}, sort string, limit, offset int) ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindWithFilter(filters, sort, limit, offset)
}

func (s *stockTransactionDomain) CreateReversalStockTransaction(stockTransaction model.StockTransaction) model.StockTransaction {
	var (
		transactionType string = "IN"
	)

	if stockTransaction.Type == "IN" {
		transactionType = "OUT"
	}

	return model.StockTransaction{
		BranchOriginID:      stockTransaction.BranchOriginID,
		BranchDestinationID: stockTransaction.BranchDestinationID,
		ItemID:              stockTransaction.ItemID,
		Type:                transactionType,
		IssuerID:            stockTransaction.IssuerID,
		Quantity:            stockTransaction.Quantity,
		Cost:                stockTransaction.Cost,
		Unit:                stockTransaction.Unit,
		ReferenceType:       stockTransaction.ReferenceType,
		Reference:           stockTransaction.Reference,
	}
}

func (s *stockTransactionDomain) InvalidateStockTransaction(ctx context.Context, filter []map[string]interface{}, userID string) (errW *error_wrapper.ErrorWrapper) {
	stockTransactions, errW := s.FindWithFilter(filter, "created_at DESC", 0, 0)

	if errW != nil {
		return
	}
	fmt.Printf("Trying to invalidate %d stock transactions", len(stockTransactions))
	for _, stockTransaction := range stockTransactions {
		now := time.Now()
		stockTransaction.DeletedAt = &now
		stockTransaction.DeletedBy = &userID
		errW = s.stockTransactionResource.Update(stockTransaction.UUID, stockTransaction)
		if errW != nil {
			return
		}
	}
	return
}
