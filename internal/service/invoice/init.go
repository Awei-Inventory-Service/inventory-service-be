package invoice

import "github.com/inventory-service/internal/repository/invoice"

func NewInvoiceService(invoiceRepository invoice.InvoiceRepository) InvoiceService {
	return &invoiceService{invoiceRepository: invoiceRepository}
}