package purchase

import "github.com/inventory-service/internal/model"

func (p *purchaseRepository) Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) error {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}

	result := p.db.Create(&purchase)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *purchaseRepository) FindAll() ([]model.Purchase, error) {
	var purchases []model.Purchase
	result := p.db.Find(&purchases)
	if result.Error != nil {
		return nil, result.Error
	}

	return purchases, nil
}

func (p *purchaseRepository) FindByID(id string) (*model.Purchase, error) {
	var purchase model.Purchase
	result := p.db.Where("uuid = ?", id).First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}

	return &purchase, nil
}

func (p *purchaseRepository) Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) error {
	purchase := model.Purchase{
		SupplierID:   supplierId,
		BranchID:     branchId,
		ItemID:       itemId,
		Quantity:     quantity,
		PurchaseCost: purchaseCost,
	}

	result := p.db.Where("uuid = ?", id).Updates(&purchase)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *purchaseRepository) Delete(id string) error {
	result := p.db.Where("uuid = ?", id).Delete(&model.Purchase{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
