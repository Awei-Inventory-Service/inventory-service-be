package invoice

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *invoiceRepository) Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
	invoice := model.Invoice{
		FileURL:     fileUrl,
		Amount:      amount,
		AmountOwed:  amountOwed,
		Notes:       notes,
		InvoiceDate: invoiceDate,
	}

	result := i.db.Create(&invoice)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (i *invoiceRepository) FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper) {
	var invoices []model.Invoice
	result := i.db.Find(&invoices)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return invoices, nil
}

func (i *invoiceRepository) FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper) {
	var invoice model.Invoice
	result := i.db.Where("uuid = ?", id).First(&invoice)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &invoice, nil
}

func (i *invoiceRepository) Update(id string, fileUrl, notes, status, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
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

	result := i.db.Model(&model.Invoice{}).
		Where("uuid = ?", id).
		Updates(updates)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (i *invoiceRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := i.db.Where("uuid = ?", id).Delete(&model.Invoice{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
