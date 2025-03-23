package invoice

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/invoice"
	"github.com/inventory-service/lib/error_wrapper"
)

type InvoiceDomain interface {
	Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper)
	Update(id string, fileUrl, notes, status, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type invoiceDomain struct {
	invoiceResource invoice.InvoiceResource
}
