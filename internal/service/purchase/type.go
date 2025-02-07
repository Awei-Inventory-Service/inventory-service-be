package purchase

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	"github.com/inventory-service/internal/repository/purchase"
	"github.com/inventory-service/internal/repository/supplier"
	"github.com/inventory-service/lib/error_wrapper"
)

type PurchaseService interface {
	Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseService struct {
	purchaseRepository purchase.PurchaseRepository
	supplierRepository supplier.SupplierRepository
	branchRepository   branch.BranchRepository
	itemRepository     item.ItemRepository
}
