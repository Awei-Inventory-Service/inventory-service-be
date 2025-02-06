package invoice

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *invoiceService) Create(fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
	err := i.invoiceRepository.Create(fileUrl, notes, invoiceDate, amount, amountOwed)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoiceService) FindAll() ([]model.Invoice, *error_wrapper.ErrorWrapper) {
	invoices, err := i.invoiceRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *invoiceService) FindByID(id string) (*model.Invoice, *error_wrapper.ErrorWrapper) {
	invoice, err := i.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (i *invoiceService) Update(id string, fileUrl, notes, invoiceDate string, amount, amountOwed float64) *error_wrapper.ErrorWrapper {
	status := "unpaid"
	if amountOwed < amount {
		status = "paid"
	}

	err := i.invoiceRepository.Update(id, fileUrl, notes, status, invoiceDate, amount, amountOwed)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoiceService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := i.invoiceRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
