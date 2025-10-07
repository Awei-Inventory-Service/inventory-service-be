package branch_product

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (b *branchProductHandler) GetBranchProductList(c *gin.Context) {
	var (
		req struct {
			Filters []dto.Filter `json:"filters"`
			Orders  []dto.Order  `json:"orders"`
			Limit   int          `json:"limit"`
			Offset  int          `json:"offset"`
		}
		errW                     *error_wrapper.ErrorWrapper
		getBranchProductResponse []dto.GetBranchProductResponse
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, getBranchProductResponse, errW)
	}()

	if err := c.ShouldBindJSON(&req); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	getBranchProductResponse, errW = b.branchProductUsecase.Get(c, req.Filters, req.Orders, req.Limit, req.Offset)

}
