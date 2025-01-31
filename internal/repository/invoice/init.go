package invoice

import "gorm.io/gorm"

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}