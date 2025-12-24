package sales_product_resource

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesProductResource) Create(ctx context.Context, payload model.SalesProduct) (newSalesProduct model.SalesProduct, errW *error_wrapper.ErrorWrapper) {
	result := s.db.Create(payload)
	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
		return
	}
	return payload, nil
}

func (s *salesProductResource) Update(ctx context.Context, payload model.SalesProduct) (newSalesProduct model.SalesProduct, errW *error_wrapper.ErrorWrapper) {
	result := s.db.Where("uuid = ? ", payload.UUID).Updates(payload)
	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return
	}
	return payload, nil
}

func (s *salesProductResource) Delete(ctx context.Context, filter model.SalesProduct) (errW *error_wrapper.ErrorWrapper) {
	query := s.db.WithContext(ctx).Model(&model.SalesProduct{})

	if filter.UUID != "" {
		query = query.Where("uuid = ?", filter.UUID)
	}

	if filter.SalesID != "" {
		query = query.Where("sales_id = ? ", filter.SalesID)
	}

	result := query.Delete(&model.SalesProduct{})
	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
		return
	}
	return
}

func (s *salesProductResource) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (salesProducts []model.SalesProduct, errW *error_wrapper.ErrorWrapper) {
	db := s.db.Model(&model.SalesProduct{})

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
		Preload("Product").
		Find(&salesProducts)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return
	}
	return
}
