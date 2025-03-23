package supplier

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/dto"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

func (s *supplierController) GetSuppliers(c *gin.Context) {
	var (
		errW      *error_wrapper.ErrorWrapper
		suppliers []model.Supplier
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, suppliers, errW)
	}()

	suppliers, errW = s.supplierService.FindAll()
	if errW != nil {
		return
	}
}

func (s *supplierController) GetSupplier(c *gin.Context) {
	var (
		errW     *error_wrapper.ErrorWrapper
		supplier *model.Supplier
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, supplier, errW)
	}()

	id := c.Param("id")
	supplier, errW = s.supplierService.FindByID(id)
	if errW != nil {
		return
	}
}

func (s *supplierController) CreateSupplier(c *gin.Context) {
	var (
		errW                  *error_wrapper.ErrorWrapper
		createSupplierRequest dto.CreateSupplierRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createSupplierRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = s.supplierService.Create(createSupplierRequest.Name, createSupplierRequest.PhoneNumber, createSupplierRequest.Address, createSupplierRequest.PICName)
	if errW != nil {
		return
	}
}

func (s *supplierController) UpdateSupplier(c *gin.Context) {
	var (
		errW                  *error_wrapper.ErrorWrapper
		updateSupplierRequest dto.UpdateSupplierRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateSupplierRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = s.supplierService.Update(id, updateSupplierRequest.Name, updateSupplierRequest.PhoneNumber, updateSupplierRequest.Address, updateSupplierRequest.PICName)
	if errW != nil {
		return
	}

}

func (s *supplierController) DeleteSupplier(c *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	errW = s.supplierService.Delete(id)
	if errW != nil {
		return
	}
}
