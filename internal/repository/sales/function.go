package sales

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *salesRepository) Create(sale model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper) {
	result := s.db.Create(&sale)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &sale, nil
}

func (s *salesRepository) FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper) {
	var sales []model.Sales
	result := s.db.Find(&sales)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return sales, nil
}

func (s *salesRepository) FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	var sale model.Sales
	result := s.db.Where("uuid = ?", id).First(&sale)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &sale, nil
}

func (s *salesRepository) Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&sale)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *salesRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.Sales{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
