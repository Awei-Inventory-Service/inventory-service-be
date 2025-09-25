package dto

import (
	"github.com/inventory-service/model"
)

type CreateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Unit             string                         `json:"unit" binding:"required"`
	ItemCompositions []CreateItemCompositionRequest `json:"item_compositions"`
}

type CreateItemCompositionRequest struct {
	ItemID string `json:"item_id" binding:"required"`
}

type UpdateItemCompositionRequest struct {
	ItemID string `json:"item_id" binding:"required"`
}

type UpdateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Unit             string                         `json:"unit" binding:"required"`
	ItemCompositions []UpdateItemCompositionRequest `json:"item_compositions"`
}

type GetItemCompositionResponse struct {
	ChildItemID   string `json:"child_item_id"`
	ChildItemName string `json:"child_item_name"`
	Unit          string `json:"unit"`
}

type GetItemsResponse struct {
	UUID              string                       `json:"uuid"`
	Name              string                       `json:"name"`
	Category          model.ItemCategory           `json:"category"`
	Unit              string                       `json:"unit"`
	ChildCompositions []GetItemCompositionResponse `json:"child_compositions"`
}
