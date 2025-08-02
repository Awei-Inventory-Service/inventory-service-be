package supplier

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *supplierDomain) Create(name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}
	return s.supplierResource.Create(supplier)
}

func (s *supplierDomain) FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper) {
	return s.supplierResource.FindAll()
}

func (s *supplierDomain) FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper) {
	return s.supplierResource.FindByID(id)
}

func (s *supplierDomain) Update(id, name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper {
	supplier := model.Supplier{
		Name:        name,
		PhoneNumber: phoneNumber,
		Address:     address,
		PICName:     picName,
	}

	return s.supplierResource.Update(id, supplier)
}

func (s *supplierDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return s.supplierResource.Delete(id)
}
