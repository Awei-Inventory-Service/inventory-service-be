package purchase

import (
	"context"

	"github.com/inventory-service/dto"
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

func (p *purchaseResource) Delete(id string) (*model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchase model.Purchase

	// First, get the purchase data before deleting
	result := p.db.Where("uuid = ?", id).First(&purchase)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	// Then delete it
	result = p.db.Where("uuid = ?", id).Delete(&model.Purchase{})
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return &purchase, nil
}

func (p *purchaseResource) FindByItemID(itemID string) ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchases []model.Purchase
	result := p.db.Where("item_id = ?", itemID).Find(&purchases)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}
	return purchases, nil
}

func (p *purchaseResource) FindByBranchAndItem(branchID, itemID string, offset, limit int) ([]model.Purchase, *error_wrapper.ErrorWrapper) {
	var purchases []model.Purchase
	result := p.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Order("created_at DESC").Offset(offset).Limit(limit).Preload("Item").Find(&purchases)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}
	return purchases, nil
}

func (p *purchaseResource) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (results []model.Purchase, errW *error_wrapper.ErrorWrapper) {
	db := p.db.Model(&model.Purchase{})

	for _, fil := range filter {
		if len(fil.Values) == 1 {
			value := fil.Values[0]

			switch fil.Wildcard {
			case "==":
				db = db.Where(fil.Key+" = ?", value)
			case "<":
				db = db.Where(fil.Key+" < ?", value)
			}
		} else {
			db = db.Where(fil.Key+" IN ?", fil.Values)
		}
	}

	for _, ord := range order {
		if ord.IsAsc {
			db = db.Order(ord.Key + " ASC")
		} else {
			db = db.Order(ord.Key + " DESC")
		}
	}

	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}

	result := db.WithContext(ctx).
		Preload("Branch").
		Preload("Supplier").
		Preload("Item").
		Find(&results)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return
	}

	return
}
