package purchase

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	"github.com/inventory-service/internal/repository/purchase"
	"github.com/inventory-service/internal/repository/supplier"
	// "github.com/inventory-service/lib/error_wrapper"
)

// TODO: Change error with error_wrapper.ErrorWrapper
type PurchaseService interface {
	Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) error
	FindAll() ([]model.Purchase, error)
	FindByID(id string) (*model.Purchase, error)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) error
	Delete(id string) error
}

type purchaseService struct {
	purchaseRepository purchase.PurchaseRepository
	supplierRepository supplier.SupplierRepository
	branchRepository   branch.BranchRepository
	itemRepository     item.ItemRepository
}
