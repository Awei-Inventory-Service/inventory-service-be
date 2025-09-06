package purchase

import (
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseDomain) Create(payload dto.CreatePurchaseRequest, userID string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase := model.Purchase{
		SupplierID:   payload.SupplierID,
		BranchID:     payload.BranchID,
		ItemID:       payload.ItemID,
		Quantity:     payload.Quantity,
		PurchaseCost: payload.PurchaseCost,
		Unit:         payload.Unit,
	}

	_, errW := p.branchItemResource.FindByBranchAndItem(payload.BranchID, payload.ItemID)

	if errW != nil {
		if errW.Is(model.RErrDataNotFound) {
			// If there is no stock balance, create one
			errW = p.branchItemResource.Create(model.BranchItem{
				BranchID:     payload.BranchID,
				ItemID:       payload.ItemID,
				CurrentStock: payload.Quantity,
			})

		} else {
			return nil, errW

		}
	}

	// Inserting new stock transaction

	errW = p.stockTransactionResource.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            userID,
		Quantity:            payload.Quantity,
		Cost:                payload.PurchaseCost,
		Unit:                purchase.Unit,
	})

	if errW != nil {
		return nil, errW
	}

	// currentItemStock, errW := p.item
	return p.purchaseResource.Create(payload.SupplierID, purchase)
}

func (p *purchaseDomain) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindAll()
}

func (p *purchaseDomain) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindByID(id)
}

func (p *purchaseDomain) Update(id, supplierId, branchId, itemId string, quantity float64, purchaseCost float64) *error_wrapper.ErrorWrapper {
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
