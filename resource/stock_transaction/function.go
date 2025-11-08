package stocktransaction

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockTransactionResource) Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transactions []model.StockTransaction
	result := s.db.Find(&transactions)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return transactions, nil
}

func (s *stockTransactionResource) FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transaction model.StockTransaction
	result := s.db.Where("uuid = ?", id).First(&transaction)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &transaction, nil
}

func (s *stockTransactionResource) Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.StockTransaction{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) Get(ctx context.Context, filters []dto.Filter, order []dto.Order, limit, offset int) (stockTransactions []model.StockTransaction, errW *error_wrapper.ErrorWrapper) {
	db := s.db.Model(&model.StockTransaction{})

	for _, filter := range filters {
		if len(filter.Values) == 1 {
			value := filter.Values[0]

			// Handle nil values for IS NULL queries
			if value == "nil" || value == "" {
				db = db.Where(filter.Key + " IS NULL")
				continue
			}

			switch filter.Wildcard {
			case "==":
				db = db.Where(filter.Key+" = ?", value)
			case "<":
				db = db.Where(filter.Key+" < ?", value)
			case "<=":
				db = db.Where(filter.Key+" <= ?", value)
			case ">":
				db = db.Where(filter.Key+" > ?", value)
			case ">=":
				db = db.Where(filter.Key+" >= ?", value)
			default:
				db = db.Where(filter.Key+" = ?", value)
			}
		} else {
			db = db.Where(filter.Key+" IN ?", filter.Values)
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

	result := db.WithContext(ctx).Preload("Item").Offset(offset).Find(&stockTransactions)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
		return
	}

	return
}

func (s *stockTransactionResource) FindWithFilter(filters []map[string]interface{}, sort string, limit, offset int) ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transactions []model.StockTransaction
	query := s.db

	for _, filter := range filters {
		field, okField := filter["field"].(string)
		value, okValue := filter["value"]
		if okField && okValue {
			if value == nil {
				query = query.Where(field + " IS NULL")
			} else {
				query = query.Where(field+" = ?", value)
			}
		}
	}

	if sort != "" {
		query = query.Order(sort)
	}

	if limit != 0 {
		query = query.Limit(limit)
	}

	result := query.Preload("Item").Offset(offset).Find(&transactions)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return transactions, nil
}
