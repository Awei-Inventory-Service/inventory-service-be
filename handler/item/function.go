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

	items, errW = i.itemService.FindAll()
	if errW != nil {
		return
	}

}

func (i *itemController) GetItem(c *gin.Context) {
	var (
		item *model.Item
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, item, errW)
	}()

	id := c.Param("id")

	item, errW = i.itemService.FindByID(id)
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

	errW = i.itemService.Create(
		createItemRequest.Name,
		createItemRequest.SupplierID,
		createItemRequest.Category,
		createItemRequest.Unit,
		createItemRequest.Price,
		createItemRequest.PortionSize,
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

	errW = i.itemService.Update(id, updateItemRequest.SupplierID, updateItemRequest.Name, updateItemRequest.Category, updateItemRequest.Price, updateItemRequest.Unit)
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
	errW = i.itemService.Delete(id)
	if errW != nil {
		return
	}

}
