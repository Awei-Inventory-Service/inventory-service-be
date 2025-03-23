package supplier

import "github.com/inventory-service/app/resource/supplier"

func NewSupplierDomain(supplierResource supplier.SupplierResource) SupplierDomain {
	return &supplierDomain{supplierResource: supplierResource}
}
