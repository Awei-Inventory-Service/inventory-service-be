package invoice

import "github.com/inventory-service/domain/invoice"

func NewInvoiceService(invoiceDomain invoice.InvoiceDomain) InvoiceService {
	return &invoiceService{invoiceDomain: invoiceDomain}
}
