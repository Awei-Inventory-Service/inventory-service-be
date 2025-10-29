package purchase

import (
	"context"

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

	// 1. Create the purchase record first
	createdPurchase, errW := p.purchaseResource.Create(payload.SupplierID, purchase)
	if errW != nil {
		return nil, errW
	}

	return createdPurchase, nil
}

func (p *purchaseDomain) FindAll() (payload []dto.GetPurchaseResponse, errW *error_wrapper.ErrorWrapper) {
	result, errW := p.purchaseResource.FindAll()

	if errW != nil {
		return
	}

	for _, data := range result {
		payload = append(payload, dto.GetPurchaseResponse{
			UUID:       data.UUID,
			Supplier:   data.Supplier,
			BranchID:   data.Branch.UUID,
			BranchName: data.Branch.Name,
			ItemID:     data.Item.UUID,
			ItemName:   data.Item.Name,
			Quantity:   data.Quantity,
			Unit:       data.Unit,
			Cost:       data.PurchaseCost,
		})
	}
	return
}

func (p *purchaseDomain) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (payload []dto.GetPurchaseResponse, errW *error_wrapper.ErrorWrapper) {
	result, errW := p.purchaseResource.Get(ctx, filter, order, limit, offset)

	if errW != nil {
		return
	}

	for _, data := range result {
		payload = append(payload, dto.GetPurchaseResponse{
			UUID:       data.UUID,
			Supplier:   data.Supplier,
			BranchID:   data.Branch.UUID,
			BranchName: data.Branch.Name,
			ItemID:     data.Item.UUID,
			ItemName:   data.Item.Name,
			Quantity:   data.Quantity,
			Unit:       data.Unit,
			Cost:       data.PurchaseCost,
		})
	}
	return
}

func (p *purchaseDomain) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	return p.purchaseResource.FindByID(id)
}

func (p *purchaseDomain) Update(id string, payload dto.UpdatePurchaseRequest) *error_wrapper.ErrorWrapper {
	purchase := model.Purchase{
		SupplierID:   payload.SupplierID,
		BranchID:     payload.BranchID,
		ItemID:       payload.ItemID,
		Quantity:     payload.Quantity,
		PurchaseCost: payload.PurchaseCost,
	}
	return p.purchaseResource.Update(id, purchase)
}

func (p *purchaseDomain) Delete(ctx context.Context, id, userID string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	// 1. Delete the purchase and get the deleted data
	return p.purchaseResource.Delete(id)
}
