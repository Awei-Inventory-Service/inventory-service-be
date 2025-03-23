package invoice

import "github.com/inventory-service/app/resource/invoice"

func NewInvoiceDomain(invoiceResource invoice.InvoiceResource) InvoiceDomain {
	return &invoiceDomain{invoiceResource: invoiceResource}
}
