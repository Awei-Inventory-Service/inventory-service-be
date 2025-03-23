package invoice

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/invoice"
)

type InvoiceController interface {
	GetInvoices(c *gin.Context)
	GetInvoice(c *gin.Context)
	CreateInvoice(c *gin.Context)
	UpdateInvoice(c *gin.Context)
	DeleteInvoice(c *gin.Context)
}

type invoiceController struct {
	invoiceService invoice.InvoiceService
}
