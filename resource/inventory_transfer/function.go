package inventory_transfer

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryTransferResource) Create(ctx context.Context, payload model.InventoryTransfer) (newData model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Create(&payload)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
		return
	}
	return payload, nil
}

func (i *inventoryTransferResource) Update(ctx context.Context, id string, payload model.InventoryTransfer) (updatedData model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Where("uuid = ?", id).Updates(&payload)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error)
		return
	}

	return payload, nil
}

func (i *inventoryTransferResource) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (results []model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	db := i.db.Model(&model.InventoryTransfer{})

	for _, fil := range filter {
		if len(fil.Values) == 1 {
			value := fil.Values[0]

			switch fil.Wildcard {
			case "==":
				db = db.Where(fil.Key+" = ?", value)
			case "<":
				db = db.Where(fil.Key+" < ?", value)
			}
		} else {
			db = db.Where(fil.Key+" IN ?", fil.Values)
		}
	}

	for _, ord := range order {
		if ord.IsAsc {
			db = db.Order(ord.Key + " ASC")
		} else {
			db = db.Order(ord.Key + " DESC")
		}
	}

	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}

	result := db.WithContext(ctx).
		Preload("BranchOrigin").
		Preload("BranchDestination").
		Preload("Items").
		Preload("Items.Item").
		Find(&results)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return
	}

	return
}

func (i *inventoryTransferResource) FindByID(ctx context.Context, id string) (inventoryTransfer model.InventoryTransfer, errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Where("uuid = ? ", id).Preload("Items").First(&inventoryTransfer)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error)
		return
	}
	return
}

func (i *inventoryTransferResource) UpdateStatus(ctx context.Context, id, status string) (errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Model(&model.InventoryTransfer{}).Where("uuid = ?", id).Update("status", status)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return
}

func (i *inventoryTransferResource) Delete(ctx context.Context, payload model.InventoryTransfer) (errW *error_wrapper.ErrorWrapper) {
	query := i.db.WithContext(ctx)

	if payload.UUID != "" {
		query = query.Where("uuid = ? ", payload.UUID)
	}

	if payload.BranchDestinationID != "" {
		query = query.Where("branch_destination_id = ? ", payload.BranchDestinationID)
	}

	if payload.BranchOriginID != "" {
		query = query.Where("branch_origin_id = ? ", payload.BranchOriginID)
	}

	result := query.Delete(&model.InventoryTransfer{})

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
		return
	}

	return
}
