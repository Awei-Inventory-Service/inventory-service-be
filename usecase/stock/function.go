package stock

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockService) GetStockByItemID(itemID string) (model.Stock, *error_wrapper.ErrorWrapper) {
	var quantity = 0.0

	// TODO: create type safe filter
	itemIdFilter := make(map[string]interface{})
	itemIdFilter["item_id"] = itemID

	transactions, err := s.stockTransactionDomain.FindWithFilter([]map[string]interface{}{itemIdFilter}, "")
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
		Quantity: 0.0,
	}, nil
}
