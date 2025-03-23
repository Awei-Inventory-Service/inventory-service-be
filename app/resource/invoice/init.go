package invoice

import "gorm.io/gorm"

func NewInvoiceResource(db *gorm.DB) InvoiceResource {
	return &invoiceResource{db: db}
}
