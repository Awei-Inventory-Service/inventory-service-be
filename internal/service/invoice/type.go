package invoice

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/invoice"
)

type InvoiceService interface {
	Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) error
	FindAll() ([]model.Invoice, error)
	FindByID(id string) (*model.Invoice, error)
	Update(id string, fileUrl, notes, invoiceDate string, amount, amountOwed float64) error
	Delete(id string) error
}

type invoiceService struct {
	invoiceRepository invoice.InvoiceRepository
}