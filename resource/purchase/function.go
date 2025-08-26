package purchase

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseResource) Create(supplierId string, purchase model.Purchase) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	result := p.db.Create(&purchase)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &purchase, nil
}

func (p *purchaseResource) FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchases []model.Purchase
	result := p.db.Preload("Branch").Preload("Supplier").Preload("Item").Find(&purchases)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return purchases, nil
}

func (p *purchaseResource) FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchase model.Purchase
	result := p.db.Preload("Supplier").Preload("Branch").Where("uuid = ?", id).First(&purchase)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &purchase, nil
}

func (p *purchaseResource) Update(id string, purchase model.Purchase) *error_wrapper.ErrorWrapper {
	result := p.db.Where("uuid = ?", id).Updates(&purchase)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (p *purchaseResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := p.db.Where("uuid = ?", id).Delete(&model.Purchase{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
