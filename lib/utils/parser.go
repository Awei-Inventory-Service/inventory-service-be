package utils

import (
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func ParseItemCategory(category string) (model.ItemCategory, *error_wrapper.ErrorWrapper) {
	switch category {
	case "raw":
		return model.ItemCategoryRaw, nil
	case "half-processed":
		return model.ItemCategoryHalfProcessed, nil
	case "finished":
		return model.ItemCategoryProcessed, nil
	case "other":
		return model.ItemCategoryOther, nil
	}

	return "", error_wrapper.New(model.UErrInvalidItemCategory, fmt.Sprintf("Category %s is invalid", category))
}
