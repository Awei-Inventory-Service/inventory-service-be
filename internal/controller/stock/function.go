package stock

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

func (s *stockController) GetStockByItemID(c *gin.Context) {
	var (
		errW  *error_wrapper.ErrorWrapper
		stock model.Stock
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, stock, errW)
	}()

	id := c.Param("id")

	stock, errW = s.stockService.GetStockByItemID(id)

	if errW != nil {
		return
	}

}
