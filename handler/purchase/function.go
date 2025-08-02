package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseController) GetPurchases(c *gin.Context) {
	var (
		purchases []model.Purchase
		errW      *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, purchases, errW)
	}()

	purchases, errW = p.purchaseService.FindAll()
	if errW != nil {
		return
	}
}

func (p *purchaseController) GetPurchase(c *gin.Context) {
	var (
		purchase *model.Purchase
		errW     *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, purchase, errW)
	}()

	id := c.Param("id")
	purchase, errW = p.purchaseService.FindByID(id)
	if errW != nil {
		return
	}
}

func (p *purchaseController) CreatePurchase(c *gin.Context) {
	var (
		createPurchaseRequest dto.CreatePurchaseRequest
		errW                  *error_wrapper.ErrorWrapper
	)

	if err := c.ShouldBindJSON(&createPurchaseRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		response_wrapper.New(&c.Writer, c, false, nil, errW)
		return
	}

	errW = p.purchaseService.Create(
		c,
		createPurchaseRequest.SupplierID,
		createPurchaseRequest.BranchID,
		createPurchaseRequest.ItemID,
		createPurchaseRequest.Quantity,
		createPurchaseRequest.PurchaseCost,
	)
	if errW != nil {
		response_wrapper.New(&c.Writer, c, false, nil, errW)
		return
	}

	response_wrapper.New(&c.Writer, c, true, nil, errW)
}

func (p *purchaseController) UpdatePurchase(c *gin.Context) {
	id := c.Param("id")
	var (
		updatePurchaseRequest dto.UpdatePurchaseRequest
		errW                  *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&updatePurchaseRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = p.purchaseService.Update(
		id,
		updatePurchaseRequest.SupplierID,
		updatePurchaseRequest.BranchID,
		updatePurchaseRequest.ItemID,
		updatePurchaseRequest.Quantity,
		updatePurchaseRequest.PurchaseCost,
	)
	if errW != nil {
		return
	}

}

func (p *purchaseController) DeletePurchase(c *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	errW = p.purchaseService.Delete(id)
	if errW != nil {
		return
	}

}
