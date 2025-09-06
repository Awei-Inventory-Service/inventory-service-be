package item

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemController) GetItems(c *gin.Context) {
	var (
		items []model.Item
		errW  *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, items, errW)
	}()

	items, errW = i.itemUsecase.FindAll()
	if errW != nil {
		return
	}

}

func (i *itemController) GetItem(c *gin.Context) {
	var (
		item *dto.GetItemsResponse
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, item, errW)
	}()

	id := c.Param("id")

	item, errW = i.itemUsecase.FindByID(c, id)
	if errW != nil {
		return
	}

}

func (i *itemController) CreateItem(c *gin.Context) {
	var (
		createItemRequest dto.CreateItemRequest
		errW              *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createItemRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	// Apply default value if not set (i.e., 0 means omitted)
	if createItemRequest.PortionSize == 0 {
		createItemRequest.PortionSize = 1.0
	}

	errW = i.itemUsecase.Create(
		c,
		createItemRequest,
	)
}

func (i *itemController) UpdateItem(c *gin.Context) {
	var (
		updateItemRequest dto.UpdateItemRequest
		errW              *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateItemRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.itemUsecase.Update(c, updateItemRequest, id)
	if errW != nil {
		return
	}

}

func (i *itemController) DeleteItem(c *gin.Context) {
	var errW *error_wrapper.ErrorWrapper

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	id := c.Param("id")
	errW = i.itemUsecase.Delete(id)
	if errW != nil {
		return
	}

}
