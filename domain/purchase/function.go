package purchase

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseDomain) Create(supplierId, branchId, itemId, userId string, quantity int, purchaseCost float64) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}

	_, errW := p.stockBalanceResource.FindByBranchAndItem(branchId, itemId)

	if errW != nil {
		if errW.Is(model.RErrDataNotFound) {
			// If there is no stock balance, create one
			errW = p.stockBalanceResource.Create(model.StockBalance{
				BranchID:     branchId,
				ItemID:       itemId,
				CurrentStock: quantity,
			})
		} else {
			return nil, errW

		}
	}

	// Inserting new stock transaction

	errW = p.stockTransactionResource.Create(model.StockTransaction{
		BranchOriginID:      branchId,
		BranchDestinationID: branchId,
		ItemID:              itemId,
		Type:                "IN",
		IssuerID:            userId,
		Quantity:            quantity,
		Cost:                purchaseCost,
	})

	if errW != nil {
		return nil, errW
	}

	// currentItemStock, errW := p.item
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
