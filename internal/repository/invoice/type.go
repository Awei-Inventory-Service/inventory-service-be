package invoice

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) error
	FindAll() ([]model.Invoice, error)
	FindByID(id string) (*model.Invoice, error)
	Update(id string, fileUrl, notes, status, invoiceDate string, amount, amountOwed float64) error
	Delete(id string) error
}

type invoiceRepository struct {
	db *gorm.DB
}