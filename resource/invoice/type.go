package invoice

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type InvoiceResource interface {
	Create(invoice model.Invoice) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper)
	Update(id string, updates map[string]interface{}) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type invoiceResource struct {
	db *gorm.DB
}
