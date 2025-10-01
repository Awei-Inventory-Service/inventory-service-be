package production

import (
	"context"
	"errors"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
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

func (p *productionResource) Get(ctx context.Context, filter dto.GetProductionFilter) ([]model.Production, *error_wrapper.ErrorWrapper) {
	var productions []model.Production
	query := p.db.Model(&model.Production{})

	if filter.FinalItemID != "" {
		query = query.Where("final_item_id = ?", filter.FinalItemID)
	}

	if filter.BranchID != "" {
		query = query.Where("branch_id = ?", filter.BranchID)
	}

	result := query.WithContext(ctx).Preload("FinalItem").Preload("SourceItems.SourceItem").Preload("Branch").Find(&productions)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Production not found")
		}
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
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
