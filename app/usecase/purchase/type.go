package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/domain/branch"
	"github.com/inventory-service/app/domain/item"
	itempurchasechain "github.com/inventory-service/app/domain/item_purchase_chain"
	"github.com/inventory-service/app/domain/purchase"
	"github.com/inventory-service/app/domain/supplier"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	// "github.com/inventory-service/lib/error_wrapper"
)

// TODO: Change error with error_wrapper.ErrorWrapper
type PurchaseService interface {
	Create(c *gin.Context, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseService struct {
	purchaseDomain          purchase.PurchaseDomain
	supplierDomain          supplier.SupplierDomain
	branchDomain            branch.BranchDomain
	itemDomain              item.ItemDomain
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain
}
