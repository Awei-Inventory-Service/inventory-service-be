package purchase

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (p *purchaseRepository) Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}

	result := p.db.Create(&purchase)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &purchase, nil
}

func (p *purchaseRepository) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchases []model.Purchase
	result := p.db.Find(&purchases)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return purchases, nil
}

func (p *purchaseRepository) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchase model.Purchase
	result := p.db.Where("uuid = ?", id).First(&purchase)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &purchase, nil
}

func (p *purchaseRepository) Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}

	result := p.db.Where("uuid = ?", id).Updates(&purchase)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (p *purchaseRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := p.db.Where("uuid = ?", id).Delete(&model.Purchase{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
