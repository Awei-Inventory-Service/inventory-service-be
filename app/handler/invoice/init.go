package invoice

import "github.com/inventory-service/app/usecase/invoice"

func NewInvoiceController(invoiceService invoice.InvoiceService) InvoiceController {
	return &invoiceController{
		invoiceService: invoiceService,
	}
}
