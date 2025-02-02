package invoice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
)

func (i *invoiceController) GetInvoices(c *gin.Context) {
	invoices, err := i.invoiceService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

func (i *invoiceController) GetInvoice(c *gin.Context) {
	id := c.Param("id")
	invoice, err := i.invoiceService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (i *invoiceController) CreateInvoice(c *gin.Context) {
	var CreateInvoiceRequest dto.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&CreateInvoiceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.invoiceService.Create(CreateInvoiceRequest.FileURL, CreateInvoiceRequest.Notes, CreateInvoiceRequest.InvoiceDate, CreateInvoiceRequest.Amount, CreateInvoiceRequest.AmountOwed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice created successfully"})
}

func (i *invoiceController) UpdateInvoice(c *gin.Context) {
	id := c.Param("id")
	var updateInvoiceRequest dto.UpdateInvoiceRequest
	if err := c.ShouldBindJSON(&updateInvoiceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.invoiceService.Update(id, updateInvoiceRequest.FileURL, updateInvoiceRequest.Notes, updateInvoiceRequest.InvoiceDate, updateInvoiceRequest.Amount, updateInvoiceRequest.AmountOwed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
}

func (i *invoiceController) DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	err := i.invoiceService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}