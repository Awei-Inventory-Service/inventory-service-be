package sales

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesResource) Create(sale model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper) {
	result := s.db.Create(&sale)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &sale, nil
}

func (s *salesResource) FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper) {
	var sales []model.Sales
	result := s.db.Find(&sales)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return sales, nil
}

func (s *salesResource) FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	var sale model.Sales
	result := s.db.Where("uuid = ?", id).First(&sale)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &sale, nil
}

func (s *salesResource) Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&sale)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *salesResource) Delete(ctx context.Context, id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	var (
		sales model.Sales
	)
	result := s.db.WithContext(ctx).Where("uuid = ?", id).First(&sales)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	result = s.db.WithContext(ctx).Where("uuid = ?", id).Delete(&model.Sales{})
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return &sales, nil
}

func (s *salesResource) FindGroupedByDate(ctx context.Context) ([]model.Sales, *error_wrapper.ErrorWrapper) {
	var sales []model.Sales
	result := s.db.WithContext(ctx).
		Preload("BranchProduct.Branch").
		Preload("BranchProduct.Product").
		Order("transaction_date DESC").
		Find(&sales)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return sales, nil
}

func (s *salesResource) FindGroupedByDateAndBranch(ctx context.Context) ([]model.Sales, *error_wrapper.ErrorWrapper) {
	var sales []model.Sales
	result := s.db.WithContext(ctx).
		Preload("BranchProduct.Branch").
		Preload("BranchProduct.Product").
		Order("transaction_date DESC, branch_product_id").
		Find(&sales)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return sales, nil
}
