package dto

import (
	"github.com/inventory-service/model"
)

type CreateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Price            float64                        `json:"price"`
	Unit             string                         `json:"unit" binding:"required"`
	PortionSize      float64                        `json:"portion_size"`
	ItemCompositions []CreateItemCompositionRequest `json:"item_compositions"`
}

type CreateItemCompositionRequest struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
	Notes  string  `json:"notes"`
}

type UpdateItemCompositionRequest struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
	Notes  string  `json:"notes"`
}

type UpdateItemRequest struct {
	Name             string                         `json:"name" binding:"required"`
	SupplierID       string                         `json:"supplier_id"`
	Category         string                         `json:"category" binding:"required"`
	Price            float64                        `json:"price"`
	Unit             string                         `json:"unit" binding:"required"`
	PortionSize      float64                        `json:"portion_size"`
	ItemCompositions []UpdateItemCompositionRequest `json:"item_compositions"`
}

type GetItemCompositionResponse struct {
	UUID        string  `json:"uuid"`
	ChildItemID string  `json:"child_item_id"`
	Ratio       float64 `json:"ratio"`
	Notes       string  `json:"notes"`
}

type GetItemsResponse struct {
	UUID              string                       `json:"uuid"`
	Name              string                       `json:"name"`
	Category          model.ItemCategory           `json:"category"`
	Unit              string                       `json:"unit"`
	PortionSize       float64                      `json:"portion_size"`
	ChildCompositions []GetItemCompositionResponse `json:"child_compositions"`
}
