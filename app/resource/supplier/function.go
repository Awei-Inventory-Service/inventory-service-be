package supplier

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *supplierResource) Create(supplier model.Supplier) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&supplier)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}
	return nil
}

func (s *supplierResource) FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper) {
	var suppliers []model.Supplier
	result := s.db.Find(&suppliers)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return suppliers, nil
}

func (s *supplierResource) FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper) {
	var supplier model.Supplier
	result := s.db.Where("uuid = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &supplier, nil
}

func (s *supplierResource) Update(id string, supplier model.Supplier) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&supplier)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *supplierResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.Supplier{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
