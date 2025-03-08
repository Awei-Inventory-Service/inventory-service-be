package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	itempurchasechain "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/purchase"
	"github.com/inventory-service/internal/repository/supplier"
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
	purchaseRepository          purchase.PurchaseRepository
	supplierRepository          supplier.SupplierRepository
	branchRepository            branch.BranchRepository
	itemRepository              item.ItemRepository
	itemPurchaseChainRepository itempurchasechain.ItemPurchaseChainRepository
}
