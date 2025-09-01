package itemcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemCompositionResource) Create(ctx context.Context, itemComposition model.ItemComposition) (errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Create(&itemComposition)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
	}

	return nil
}

func (i *itemCompositionResource) DeleteByItemID(ctx context.Context, itemID string) (errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Where("parent_item_id = ?", itemID).Delete(&model.ItemComposition{})

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
	}

	return nil
}
