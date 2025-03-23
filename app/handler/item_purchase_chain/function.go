package itempurchasechain

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/dto"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

// Akan dipanggil tiap kali ada purchase baru. Jadi 1 purchase -> 1 item purchase chain
func (i *itemPurchaseChainController) Create(ctx *gin.Context) {
	var (
		itemPurchaseChain dto.CreateItemPurchaseChainRequest
		errW              *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	companyId := ctx.GetHeader("company_id")

	if companyId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "Company id is required")
		return
	}

	if err := ctx.ShouldBindJSON(&itemPurchaseChain); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.itemPurchaseChainService.Create(ctx, itemPurchaseChain.ItemID, itemPurchaseChain.BranchID, itemPurchaseChain.PurchaseID)

	if errW != nil {
		return
	}
}
