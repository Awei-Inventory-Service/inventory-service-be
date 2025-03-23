package invoice

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *invoiceResource) Create(invoice model.Invoice) *error_wrapper.ErrorWrapper {
	result := i.db.Create(&invoice)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (i *invoiceResource) FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper) {
	var invoices []model.Invoice
	result := i.db.Find(&invoices)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return invoices, nil
}

func (i *invoiceResource) FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper) {
	var invoice model.Invoice
	result := i.db.Where("uuid = ?", id).First(&invoice)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &invoice, nil
}

func (i *invoiceResource) Update(id string, updates map[string]interface{}) *error_wrapper.ErrorWrapper {
	result := i.db.Model(&model.Invoice{}).
		Where("uuid = ?", id).
		Updates(updates)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (i *invoiceResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := i.db.Where("uuid = ?", id).Delete(&model.Invoice{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
