package invoice

import (
	"github.com/inventory-service/domain/invoice"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type InvoiceService interface {
	Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper)
	Update(id string, fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type invoiceService struct {
	invoiceDomain invoice.InvoiceDomain
}
