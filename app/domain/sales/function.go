package sales

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *salesDomain) Create(sale model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.Create(sale)
}

func (s *salesDomain) FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindAll()
}

func (s *salesDomain) FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindByID(id)
}

func (s *salesDomain) Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper {
	return s.salesResource.Update(id, sale)
}

func (s *salesDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return s.salesResource.Delete(id)
}
