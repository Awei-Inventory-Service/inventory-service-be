package supplier

import (
	"github.com/inventory-service/internal/model"
)

func (s *supplierRepository) Create(name, phoneNumber, address, picName string) error {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}

	result := s.db.Create(&supplier)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *supplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	result := s.db.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

func (s *supplierRepository) FindByID(id string) (*model.Supplier, error) {
	var supplier model.Supplier
	result := s.db.Where("uuid = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, result.Error
	}

	return &supplier, nil
}

func (s *supplierRepository) Update(id, name, phoneNumber, address, picName string) error {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}

	result := s.db.Where("uuid = ?", id).Updates(&supplier)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *supplierRepository) Delete(id string) error {
	result := s.db.Where("uuid = ?", id).Delete(&model.Supplier{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
