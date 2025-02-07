package supplier

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *supplierService) Create(name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	err := s.supplierRepository.Create(name, phoneNumber, address, picName)
	if err != nil {
		return err
	}

	return nil
}

func (s *supplierService) FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper) {
	suppliers, err := s.supplierRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (s *supplierService) FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper) {
	supplier, err := s.supplierRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *supplierService) Update(id, name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	err := s.supplierRepository.Update(id, name, phoneNumber, address, picName)
	if err != nil {
		return err
	}

	return nil
}

func (s *supplierService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := s.supplierRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
