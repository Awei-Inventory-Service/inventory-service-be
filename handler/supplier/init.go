package supplier

import "github.com/inventory-service/usecase/supplier"

func NewSupplierController(supplierService supplier.SupplierService) SupplierController {
	return &supplierController{
		supplierService: supplierService,
	}
}
