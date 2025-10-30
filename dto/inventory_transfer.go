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

type UpdateInventoryTransferRequest struct {
	BranchDestinationID string                               `json:"branch_destination_id" binding:"required"`
	BranchOriginID      string                               `json:"branch_origin_id" binding:"required"`
	IssuerID            string                               `json:"issuer_id"`
	TransferDate        string                               `json:"transfer_date" binding:"required"`
	Status              string                               `json:"status" binding:"required"`
	Remarks             string                               `json:"remarks"`
	Items               []UpdateInventoryTransferItemRequest `json:"items" binding:"required"`
}

type UpdateInventoryTransferItemRequest struct {
	ItemID   string  `json:"item_id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
	Unit     string  `json:"unit" binding:"required"`
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

type GetInventoryTransferListResponse struct {
	InventoryTansfers []InventoryTransferResponse `json:"inventory_transfers"`
}
type InventoryTransferResponse struct {
	UUID                  string                             `json:"uuid"`
	BranchOriginName      string                             `json:"branch_origin_name"`
	BranchOriginID        string                             `json:"branch_origin_id"`
	BranchDestinationName string                             `json:"branch_destination_name"`
	BranchDestinationID   string                             `json:"branch_destination_id"`
	Items                 []GetInventoryTransferItemResponse `json:"items"`
	TransferDate          string                             `json:"transfer_date"`
	Status                string                             `json:"status"`
}

type GetInventoryTransferItemResponse struct {
	UUID         string  `json:"uuid"`
	ItemID       string  `json:"item_id"`
	ItemName     string  `json:"item_name"`
	ItemQuantity float64 `json:"item_quantity"`
	ItemUnit     string  `json:"item_unit"`
}

type DeleteInventoryTransferRequest struct {
	ID       string `json:"id"`
	BranchID string `json:"branch_id"`
	UserID   string `json:"user_id"`
}

func (u UpdateInventoryTransferStatus) ValidateStatus() bool {
	if u.Status != constant.TRANSFER_STATUS_CANCELLED && u.Status != constant.TRANSFER_STATUS_IN_PROGRESS && u.Status != constant.TRANSFER_STATUS_COMPLETED {
		return false
	}

	return true
}
