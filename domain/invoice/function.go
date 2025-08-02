package invoice

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *invoiceDomain) Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
	invoice := model.Invoice{
		FileURL:     fileUrl,
		Amount:      amount,
		AmountOwed:  amountOwed,
		Notes:       notes,
		InvoiceDate: invoiceDate,
	}

	return i.invoiceResource.Create(invoice)
}

func (i *invoiceDomain) FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper) {
	return i.invoiceResource.FindAll()
}

func (i *invoiceDomain) FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper) {
	return i.invoiceResource.FindByID(id)
}

func (i *invoiceDomain) Update(id string, fileUrl, notes, status, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
	updates := map[string]interface{}{}

	if fileUrl != "" {
		updates["file_url"] = fileUrl
	}
	if amount != 0 {
		updates["amount"] = amount
	}
	if amountOwed != 0 {
		updates["amount_owed"] = amountOwed
	}
	if notes != "" {
		updates["notes"] = notes
	}
	if status != "" {
		updates["status"] = status
	}
	if invoiceDate != "" {
		updates["invoice_date"] = invoiceDate
	}

	if len(updates) == 0 {
		return nil
	}

	return i.invoiceResource.Update(id, updates)
}

func (i *invoiceDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return i.invoiceResource.Delete(id)
}
