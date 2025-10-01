package productionitem

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productionItemResource) Create(ctx context.Context, productionItem model.ProductionItem) (*model.ProductionItem, *error_wrapper.ErrorWrapper) {
	result := p.db.WithContext(ctx).Create(&productionItem)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &productionItem, nil
}

func (p *productionItemResource) Get(ctx context.Context, filter model.ProductionItem) ([]model.ProductionItem, *error_wrapper.ErrorWrapper) {
	var productionItems []model.ProductionItem
	query := p.db.WithContext(ctx).Model(&model.ProductionItem{})

	if filter.UUID != "" {
		query = query.Where("uuid = ?", filter.UUID)
	}

	if filter.ProductionID != "" {
		query = query.Where("production_id = ?", filter.ProductionID)
	}

	if filter.SourceItemID != "" {
		query = query.Where("source_item_id = ?", filter.SourceItemID)
	}

	result := query.Find(&productionItems)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return productionItems, nil
}

func (p *productionItemResource) Delete(ctx context.Context, filter model.ProductionItem) *error_wrapper.ErrorWrapper {
	query := p.db.WithContext(ctx).Model(&model.ProductionItem{})

	if filter.UUID != "" {
		query = query.Where("uuid = ?", filter.UUID)
	}

	if filter.ProductionID != "" {
		query = query.Where("production_id = ?", filter.ProductionID)
	}

	result := query.Delete(&model.ProductionItem{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
