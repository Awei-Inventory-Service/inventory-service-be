package invoice

import "github.com/inventory-service/internal/service/invoice"

func NewInvoiceController(invoiceService invoice.InvoiceService) InvoiceController {
	return &invoiceController{
		invoiceService: invoiceService,
	}
}