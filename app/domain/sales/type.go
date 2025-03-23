package sales

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/sales"
	"github.com/inventory-service/lib/error_wrapper"
)

type SalesDomain interface {
	Create(sale model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper)
	Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type salesDomain struct {
	salesResource sales.SalesResource
}
