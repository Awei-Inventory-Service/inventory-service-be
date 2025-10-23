package dto

import "github.com/inventory-service/constant"

type CreateInventoryTransferRequest struct {
	BranchDestinationID string                               `json:"branch_destination_id" binding:"required"`
	BranchOriginID      string                               `json:"branch_origin_id" binding:"required"`
	IssuerID            string                               `json:"issuer_id"`
	TransferDate        string                               `json:"transfer_date" binding:"required"`
	Remarks             string                               `json:"remarks"`
	Items               []CreateInventoryTransferItemRequest `json:"items" binding:"required"`
}

type CreateInventoryTransferItemRequest struct {
	ItemID   string  `json:"item_id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
	Unit     string  `json:"unit" binding:"required"`
}

type UpdateInventoryTransferStatus struct {
	InventoryTransferID string `json:"inventory_transfer_id" binding:"required"`
	Status              string `json:"status"`
}

func (u UpdateInventoryTransferStatus) ValidateStatus() bool {
	if u.Status != constant.TRANSFER_STATUS_CANCELLED && u.Status != constant.TRANSFER_STATUS_IN_PROGRESS && u.Status != constant.TRANSFER_STATUS_COMPLETED {
		return false
	}

	return true
}
