package purchase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (p *purchaseController) Get(c *gin.Context) {
	var (
		purchases []dto.GetPurchaseResponse
		errW      *error_wrapper.ErrorWrapper
		payload   dto.GetListRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, purchases, errW)
	}()

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	purchases, errW = p.purchaseService.Get(c, payload.Filter, payload.Order, payload.Limit, payload.Offset)
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

func (p *purchaseController) Create(c *gin.Context) {
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
		createPurchaseRequest,
	)
	if errW != nil {
		response_wrapper.New(&c.Writer, c, false, nil, errW)
		return
	}

	response_wrapper.New(&c.Writer, c, true, nil, errW)
}

func (p *purchaseController) Update(c *gin.Context) {
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

	userId := c.GetHeader("user_id")

	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is required")
		return
	}
	updatePurchaseRequest.UserID = userId

	errW = p.purchaseService.Update(
		c,
		id,
		updatePurchaseRequest,
	)
	if errW != nil {
		return
	}

}

func (p *purchaseController) Delete(c *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")

	userId := c.GetHeader("user_id")
	fmt.Println("Ini user id", userId)
	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is missing on the header")
		return
	}
	errW = p.purchaseService.Delete(c, id, userId)
	if errW != nil {
		return
	}

}
