package invoice

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

func (i *invoiceController) GetInvoices(c *gin.Context) {
	var (
		invoices []model.Invoice
		errW     *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, invoices, errW)
	}()

	invoices, errW = i.invoiceService.FindAll()

	if errW != nil {
		return
	}

}

func (i *invoiceController) GetInvoice(c *gin.Context) {
	var (
		invoice *model.Invoice
		errW    *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, invoice, errW)
	}()

	id := c.Param("id")
	invoice, errW = i.invoiceService.FindByID(id)

	if errW != nil {
		return
	}

}

func (i *invoiceController) CreateInvoice(c *gin.Context) {
	var (
		createInvoiceRequest dto.CreateInvoiceRequest
		errW                 *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createInvoiceRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.invoiceService.Create(
		createInvoiceRequest.FileURL,
		createInvoiceRequest.Notes,
		createInvoiceRequest.InvoiceDate,
		createInvoiceRequest.Amount,
		createInvoiceRequest.AmountOwed,
	)

	if errW != nil {
		return
	}

}

func (i *invoiceController) UpdateInvoice(c *gin.Context) {
	var (
		updateInvoiceRequest dto.UpdateInvoiceRequest
		errW                 *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateInvoiceRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.invoiceService.Update(id, updateInvoiceRequest.FileURL, updateInvoiceRequest.Notes, updateInvoiceRequest.InvoiceDate, updateInvoiceRequest.Amount, updateInvoiceRequest.AmountOwed)
	if errW != nil {
		return
	}

}

func (i *invoiceController) DeleteInvoice(c *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	errW = i.invoiceService.Delete(id)
	if errW != nil {
		return
	}

}
