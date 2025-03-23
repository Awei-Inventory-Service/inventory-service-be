package invoice

import "github.com/inventory-service/app/domain/invoice"

func NewInvoiceService(invoiceDomain invoice.InvoiceDomain) InvoiceService {
	return &invoiceService{invoiceDomain: invoiceDomain}
}
