package item

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	db, mock, err := mocks.SetupMockDB()

	assert.NoError(t, err)

	repo := &itemResource{db: db}
	tests := []struct {
		name string
		item struct {
			name       string
			supplierId string
			category   string
			price      float64
			unit       string
		}
		mockF func(sqlmock.Sqlmock)
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success",
			item: struct {
				name       string
				supplierId string
				category   string
				price      float64
				unit       string
			}{
				name:       "Test Item",
				supplierId: "supplier1",
				category:   "Test Category",
				price:      100.0,
				unit:       "piece",
			},
			mockF: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "items" ("name","category","price","unit","created_at","updated_at","supplier_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
					WillReturnRows(sqlmock.NewRows([]string{"uuid"}).AddRow(1))
				mock.ExpectCommit()
			},
			err: nil,
		},
		{
			name: "Fail create new item",
			item: struct {
				name       string
				supplierId string
				category   string
				price      float64
				unit       string
			}{
				name:       "Test Item",
				supplierId: "supplier1",
				category:   "Test Category",
				price:      100.0,
				unit:       "piece",
			},
			mockF: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "items" ("name","category","price","unit","created_at","updated_at","supplier_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`)).
					WillReturnError(errors.New("foreign key violation: supplier does not exist"))
				mock.ExpectRollback()
			},
			err: error_wrapper.New(model.RErrPostgresCreateDocument, "Foreign key violation"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(mock)
			item := model.Item{
				Name:       tt.item.name,
				Category:   tt.item.category,
				SupplierID: tt.item.supplierId,
				Price:      tt.item.price,
				Unit:       tt.item.unit,
			}
			err := repo.Create(
				item,
			)

			if err != nil {
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
			}

		})
	}
}
func TestFindAll(t *testing.T) {
	db, mock, err := mocks.SetupMockDB()
	assert.NoError(t, err)

	repo := &itemResource{db: db}

	tests := []struct {
		name         string
		mockF        func(sqlmock.Sqlmock)
		err          *error_wrapper.ErrorWrapper
		expectedData []model.Item
	}{
		{
			name: "Success find all",
			mockF: func(mock sqlmock.Sqlmock) {

				supplierRows := sqlmock.NewRows([]string{"uuid", "name", "phone_number", "address", "pic_name", "created_at", "updated_at"}).
					AddRow("supplier-1", "Supplier 1", "123456789", "Address 1", "PIC 1", time.Now(), time.Now())

				// Mock item query
				itemRows := sqlmock.NewRows([]string{"uuid", "name", "category", "price", "unit", "created_at", "updated_at", "supplier_id"}).
					AddRow("uuid-1", "Item 1", "Category 1", 100.0, "piece", time.Now(), time.Now(), "supplier-1")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "items"`)).WillReturnRows(itemRows)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "suppliers" WHERE "suppliers"."uuid" = $1`)).
					WillReturnRows(supplierRows)
			},
			err: nil,
			expectedData: []model.Item{
				{
					UUID:     "uuid-1",
					Name:     "Item 1",
					Category: "Category 1",
					Price:    100.0,
					Unit:     "piece",
				},
			},
		},
		{
			name: "Success find all",
			mockF: func(mock sqlmock.Sqlmock) {
				supplierRows := sqlmock.NewRows([]string{"uuid", "name", "phone_number", "address", "pic_name", "created_at", "updated_at"}).
					AddRow("supplier-1", "Supplier 1", "123456789", "Address 1", "PIC 1", time.Now(), time.Now()).
					AddRow("supplier-2", "Supplier 2", "123456789", "Address 2", "PIC 2", time.Now(), time.Now())

				itemRows := sqlmock.NewRows([]string{"uuid", "name", "category", "price", "unit", "created_at", "updated_at", "supplier_id"}).
					AddRow("uuid-1", "Item 1", "Category 1", 100.0, "piece", time.Now(), time.Now(), "supplier-1").
					AddRow("uuid-2", "Item 2", "Category 2", 110.0, "piece", time.Now(), time.Now(), "supplier-2")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "items"`)).WillReturnRows(itemRows)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "suppliers" WHERE "suppliers"."uuid" IN ($1,$2`)).
					WillReturnRows(supplierRows)
			},
			err: nil,
			expectedData: []model.Item{
				{
					UUID:       "uuid-1",
					Name:       "Item 1",
					Category:   "Category 1",
					Price:      100.0,
					Unit:       "piece",
					SupplierID: "supplier-1",
					Supplier: model.Supplier{
						UUID:        "supplier-1",
						Name:        "Supplier 1",
						PhoneNumber: "123456789",
						Address:     "Addresss 1",
						PICName:     "PIC 1",
					},
				},
				{
					UUID:       "uuid-2",
					Name:       "Item 2",
					Category:   "Category 2",
					Price:      110.0,
					Unit:       "piece",
					SupplierID: "supplier-2",
					Supplier: model.Supplier{
						UUID:        "supplier-2",
						Name:        "Supplier 2",
						PhoneNumber: "123456789",
						Address:     "Addresss 2",
						PICName:     "PIC 2",
					},
				},
			},
		},
		{
			name: "Failed to find all",
			mockF: func(s sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "items"`)).WillReturnError(errors.New("Error getting all items"))
			},
			err:          error_wrapper.New(model.RErrPostgresReadDocument, "Error getting all items"),
			expectedData: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(mock)

			result, err := repo.FindAll()

			// Debugging
			fmt.Println("Result:", result)
			fmt.Println("Error:", err)

			// Ensure mock expectations were met (fix: don't assign to `err`)
			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())

			} else {
				var filteredResult, filteredExpected []model.Item
				for _, item := range result {
					filteredResult = append(filteredResult, model.Item{
						UUID:       item.UUID,
						Name:       item.Name,
						Category:   item.Category,
						Price:      item.Price,
						Unit:       item.Unit,
						SupplierID: item.SupplierID,
						Supplier:   item.Supplier,
					})
				}
				for _, item := range tt.expectedData {
					filteredExpected = append(filteredExpected, model.Item{
						UUID:       item.UUID,
						Name:       item.Name,
						Category:   item.Category,
						Price:      item.Price,
						Unit:       item.Unit,
						SupplierID: item.SupplierID,
						Supplier:   item.Supplier,
					})
				}

				assert.ElementsMatch(t, filteredExpected, filteredResult)
				assert.Nil(t, err)
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	db, mock, err := mocks.SetupMockDB()
	assert.NoError(t, err)

	repo := &itemResource{db: db}

	tests := []struct {
		name         string
		mockF        func(sqlmock.Sqlmock)
		err          *error_wrapper.ErrorWrapper
		expectedData model.Item
	}{
		{
			name: "Success find an item",
			mockF: func(s sqlmock.Sqlmock) {

				supplierRows := sqlmock.NewRows([]string{"uuid", "name", "phone_number", "address", "pic_name", "created_at", "updated_at"}).
					AddRow("supplier-1", "Supplier 1", "123456789", "Address 1", "PIC 1", time.Now(), time.Now())

				itemRows := sqlmock.NewRows([]string{"uuid", "name", "category", "price", "unit", "created_at", "updated_at", "supplier_id"}).
					AddRow("uuid-1", "Item 1", "Category 1", 100.0, "piece", time.Now(), time.Now(), "supplier-1")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "items" WHERE uuid = $1 ORDER BY "items"."uuid" LIMIT $2`)).
					WillReturnRows(itemRows)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "suppliers" WHERE "suppliers"."uuid" = $1`)).
					WillReturnRows(supplierRows)
			},
			err: nil,
			expectedData: model.Item{
				UUID:       "uuid-1",
				Name:       "Item 1",
				Category:   "Category 1",
				Price:      100.0,
				Unit:       "piece",
				SupplierID: "supplier-1",
				Supplier: model.Supplier{
					UUID:        "supplier-1",
					Name:        "Supplier 1",
					PhoneNumber: "123456789",
					Address:     "Addresss 1",
					PICName:     "PIC 1",
				},
			},
		},
		{
			name: "Failed getting an item",
			mockF: func(s sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "items" WHERE uuid = $1 ORDER BY "items"."uuid" LIMIT $2`)).
					WillReturnError(errors.New("Fail getting an item"))
			},
			err:          error_wrapper.New(model.RErrPostgresReadDocument, "Fail getting an item"),
			expectedData: model.Item{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(mock)

			result, err := repo.FindByID("uuid-1")

			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)

				assert.Equal(t, tt.expectedData.UUID, result.UUID)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := mocks.SetupMockDB()
	assert.NoError(t, err)

	repo := &itemResource{db: db}

	// Use a fixed time to avoid precision mismatches
	fixedUpdatedAt := time.Date(2025, 2, 6, 19, 5, 21, 628451000, time.FixedZone("WIB", 7*3600))

	item := model.Item{
		UUID:       "uuid-1",
		Name:       "Updated Item",
		SupplierID: "supplier1",
		Category:   "Updated Category",
		Price:      200.0,
		Unit:       "box",
		UpdatedAt:  fixedUpdatedAt, // Use fixed time
	}

	tests := []struct {
		name  string
		mockF func(sqlmock.Sqlmock)
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success",
			mockF: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "items" SET "name"=$1,"category"=$2,"price"=$3,"unit"=$4,"updated_at"=$5,"supplier_id"=$6 WHERE uuid = $7`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			err: nil,
		},
		{
			name: "Failed updating an item",
			mockF: func(s sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "items" SET "name"=$1,"category"=$2,"price"=$3,"unit"=$4,"updated_at"=$5,"supplier_id"=$6 WHERE uuid = $7`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("Error updating item"))
				mock.ExpectRollback()
			},
			err: error_wrapper.New(model.RErrPostgresUpdateDocument, "Error updating item"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(mock)

			err := repo.Update(item.UUID, item)

			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := mocks.SetupMockDB()
	assert.NoError(t, err)

	repo := &itemResource{db: db}

	tests := []struct {
		name  string
		mockF func(sqlmock.Sqlmock)
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success delete",
			mockF: func(mock sqlmock.Sqlmock) {
				// Expect the transaction to begin
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "items" WHERE uuid = $1`)).
					WithArgs("uuid-1").
					WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected

				// Expect the transaction to be committed
				mock.ExpectCommit()
			},
			err: nil,
		},
		{
			name: "Failure delete",
			mockF: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "items" WHERE uuid = $1`)).
					WithArgs("uuid-1").
					WillReturnError(fmt.Errorf("some error"))

				mock.ExpectRollback()
			},
			err: error_wrapper.New(model.RErrPostgresDeleteDocument, "some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(mock) // Apply the mock expectations

			err := repo.Delete("uuid-1")

			if tt.err != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
