package inventory_transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/constant"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryTransferDomain) Create(ctx context.Context, payload dto.CreateInventoryTransferRequest) (result model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	parsedDate, err := time.Parse("2006-01-02", payload.TransferDate)

	if err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, "Invalid date format")
		fmt.Println("ERROR", errW)
		return
	}
	result, errW = i.inventoryTransferResource.Create(ctx, model.InventoryTransfer{
		BranchOriginID:      payload.BranchOriginID,
		BranchDestinationID: payload.BranchDestinationID,
		Status:              constant.TRANSFER_STATUS_IN_PROGRESS,
		TransferDate:        parsedDate,
		Remarks:             &payload.Remarks,
		IssuerID:            payload.IssuerID,
	})

	return
}

func (i *inventoryTransferDomain) FindByID(ctx context.Context, id string) (result model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryTransferResource.FindByID(ctx, id)
}

func (i *inventoryTransferDomain) UpdateStatus(ctx context.Context, id, status string) (errW *error_wrapper.ErrorWrapper) {
	return i.inventoryTransferResource.UpdateStatus(ctx, id, status)
}
