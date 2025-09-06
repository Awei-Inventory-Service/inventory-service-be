package purchase

import (
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	branchitem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type PurchaseDomain interface {
	Create(payload dto.CreatePurchaseRequest, userId string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id, supplierId, branchId, itemId string, quantity float64, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseDomain struct {
	purchaseResource         purchase.PurchaseResource
	branchItemResource       branchitem.BranchItemResource
	stockTransactionResource stocktransaction.StockTransactionResource
}
