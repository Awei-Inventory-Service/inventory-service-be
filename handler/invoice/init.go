package invoice

import "github.com/inventory-service/usecase/invoice"

func NewInvoiceController(invoiceService invoice.InvoiceService) InvoiceController {
	return &invoiceController{
		invoiceService: invoiceService,
	}
}
