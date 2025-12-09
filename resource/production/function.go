package production

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productionResource) Create(ctx context.Context, production model.Production) (*model.Production, *error_wrapper.ErrorWrapper) {
	result := p.db.WithContext(ctx).Create(&production)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &production, nil
}

func (p *productionResource) FindAll() ([]model.Production, *error_wrapper.ErrorWrapper) {
	var productions []model.Production
	result := p.db.Find(&productions)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return productions, nil
}

func (p *productionResource) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.Production, *error_wrapper.ErrorWrapper) {
	var productions []model.Production
	db := p.db.Model(&model.Production{})

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
		Preload("FinalItem").
		Preload("SourceItems.SourceItem").
		Preload("Branch").
		Find(&productions)

	if result.Error != nil {
		errW := error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return nil, errW
	}

	return productions, nil
}

func (p *productionResource) FindByID(id string) (*model.Production, *error_wrapper.ErrorWrapper) {
	var production model.Production
	result := p.db.Where("uuid = ?", id).First(&production)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &production, nil
}

func (p *productionResource) Update(id string, production model.Production) *error_wrapper.ErrorWrapper {
	result := p.db.Where("uuid = ?", id).Updates(&production)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (p *productionResource) Delete(ctx context.Context, id string) *error_wrapper.ErrorWrapper {
	result := p.db.WithContext(ctx).Where("uuid = ?", id).Delete(&model.Production{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
