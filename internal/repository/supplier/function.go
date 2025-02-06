package supplier

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *supplierRepository) Create(name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}

	result := s.db.Create(&supplier)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}
	return nil
}

func (s *supplierRepository) FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper) {
	var suppliers []model.Supplier
	result := s.db.Find(&suppliers)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return suppliers, nil
}

func (s *supplierRepository) FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper) {
	var supplier model.Supplier
	result := s.db.Where("uuid = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &supplier, nil
}

func (s *supplierRepository) Update(id, name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}

	result := s.db.Where("uuid = ?", id).Updates(&supplier)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *supplierRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.Supplier{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
