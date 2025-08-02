package invoice

import "github.com/inventory-service/resource/invoice"

func NewInvoiceDomain(invoiceResource invoice.InvoiceResource) InvoiceDomain {
	return &invoiceDomain{invoiceResource: invoiceResource}
}
