package purchase

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseDomain) Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}
	return p.purchaseResource.Create(supplierId, purchase)
}

func (p *purchaseDomain) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindAll()
}

func (p *purchaseDomain) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindByID(id)
}

func (p *purchaseDomain) Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}
	return p.purchaseResource.Update(id, purchase)
}

func (p *purchaseDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return p.purchaseResource.Delete(id)
}
