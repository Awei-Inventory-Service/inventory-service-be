package stock

import (
	"github.com/inventory-service/internal/service/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *stockService) GetStockByItemID(itemID string) (model.Stock, *error_wrapper.ErrorWrapper) {
	var quantity int

	// TODO: create type safe filter
	itemIdFilter := make(map[string]interface{})
	itemIdFilter["item_id"] = itemID

	transactions, err := s.stockTransactionRepository.FindWithFilter([]map[string]interface{}{itemIdFilter})
	if err != nil {
		return model.Stock{}, err
	}

	// TODO: Optimize the data using concurrency (if data is large)
	for _, trx := range transactions {
		if trx.Type == "IN" {
			quantity += trx.Quantity
		} else if trx.Type == "OUT" {
			quantity -= trx.Quantity
		}
	}

	return model.Stock{
		ItemID:   itemID,
		Quantity: quantity,
	}, nil
}
