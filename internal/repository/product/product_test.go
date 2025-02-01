package product_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/product"
	"github.com/inventory-service/lib/error_wrapper"
	mock_mongodb "github.com/inventory-service/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	repo := product.NewProductRepository(mockClient, "testDB", "products")

	assert.NotNil(t, repo)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	repo := product.NewProductRepository(mockClient, "testDB", "products")

	type args struct {
		ctx         context.Context
		name        string
		ingredients []model.Ingredient
	}
	tests := []struct {
		name  string
		mockF func()
		args  args
		want  *error_wrapper.ErrorWrapper
	}{
		{
			name: "success",
			args: args{
				ctx:         context.Background(),
				name:        "Test Product",
				ingredients: []model.Ingredient{{ItemName: "Ingredient1"}},
			},
			mockF: func() {
				mockCollection.EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(nil, nil) // Simulating successful insert
			},
			want: nil,
		},
		{
			name: "failure - InsertOne error",
			args: args{
				ctx:         context.Background(),
				name:        "Test Product",
				ingredients: []model.Ingredient{{ItemName: "Ingredient1"}},
			},
			mockF: func() {
				mockCollection.EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("insert error"))
			},
			want: error_wrapper.New(model.RErrMongoDBCreateDocument, "insert error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF()
			got := repo.Create(tt.args.ctx, tt.args.name, tt.args.ingredients)
			if got == nil {
				assert.Nil(t, tt.want)
			} else {
				assert.Equal(t, tt.want.StatusCode(), got.StatusCode())
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)
	mockCursor := mock_mongodb.NewMockMongoDBCursorWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	repo := product.NewProductRepository(mockClient, "testDB", "products")

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name  string
		mockF func()
		args  args
		want  []model.Product
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Return(mockCursor, nil)

				mockCursor.EXPECT().Next(gomock.Any()).Return(true).Times(2)
				mockCursor.EXPECT().Decode(gomock.Any()).DoAndReturn(func(p *model.Product) error {
					*p = model.Product{Name: "Product1"}
					return nil
				}).Times(1)

				mockCursor.EXPECT().Decode(gomock.Any()).DoAndReturn(func(p *model.Product) error {
					*p = model.Product{Name: "Product2"}
					return nil
				}).Times(1)

				mockCursor.EXPECT().Next(gomock.Any()).Return(false)
				mockCursor.EXPECT().Err().Return(nil)
				mockCursor.EXPECT().Close(gomock.Any())
			},
			want: []model.Product{
				{Name: "Product1"},
				{Name: "Product2"},
			},
			err: nil,
		},
		{
			name: "failure - Find error",
			args: args{
				ctx: context.Background(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			want: nil,
			err:  error_wrapper.New(model.RErrMongoDBReadDocument, "database error"),
		},
		{
			name: "failure - Decode error",
			args: args{
				ctx: context.Background(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Return(mockCursor, nil)

				mockCursor.EXPECT().Next(gomock.Any()).Return(true)
				mockCursor.EXPECT().Decode(gomock.Any()).Return(errors.New("decode error"))
				mockCursor.EXPECT().Close(gomock.Any())
			},
			want: nil,
			err:  error_wrapper.New(model.RErrMongoDBReadDocument, "decode error"),
		},
		{
			name: "failure - Cursor Err",
			args: args{
				ctx: context.Background(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Return(mockCursor, nil)

				mockCursor.EXPECT().Next(gomock.Any()).Return(false)
				mockCursor.EXPECT().Err().Return(errors.New("cursor error"))
				mockCursor.EXPECT().Close(gomock.Any())
			},
			want: nil,
			err:  error_wrapper.New(model.RErrMongoDBReadDocument, "cursor error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF()
			got, err := repo.FindAll(tt.args.ctx)

			if tt.err == nil {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Nil(t, got)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)
	mockSingleResult := mock_mongodb.NewMockMongoDBSingleResultWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	repo := product.NewProductRepository(mockClient, "testDB", "products")
	type args struct {
		ctx       context.Context
		productID string
	}

	validId := primitive.NewObjectID()
	invalidID := "invalid_hex"

	ingredients := []model.Ingredient{
		{
			ItemName: "Item A",
			ItemID:   "1231",
			Quantity: 10,
			Unit:     "kg",
		},
		{
			ItemID:   "1232",
			ItemName: "Item B",
			Quantity: 12,
			Unit:     "gr",
		},
	}
	expectedProduct := model.Product{
		Name:        "Produk A",
		Ingredients: ingredients,
	}

	tests := []struct {
		name  string
		mockF func()
		args  args
		want  model.Product
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success find a data",
			args: args{
				ctx:       context.Background(),
				productID: validId.Hex(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					FindOne(gomock.Any(), bson.M{"_id": validId}).
					Return(mockSingleResult)

				mockSingleResult.EXPECT().
					Decode(gomock.Any()).
					DoAndReturn(func(p *model.Product) error {
						*p = expectedProduct
						return nil
					})
			},
			want: expectedProduct,
			err:  nil,
		},
		{
			name: "Failed on decoding id",
			args: args{
				ctx:       context.Background(),
				productID: invalidID,
			},
			mockF: func() {},
			want:  model.Product{},
			err:   error_wrapper.New(model.RErrMongoDBReadDocument, "Invalid object id"),
		},
		{
			name: "Falied when decoding product result",
			args: args{
				ctx:       context.Background(),
				productID: validId.Hex(),
			},
			mockF: func() {
				mockCollection.EXPECT().
					FindOne(gomock.Any(), bson.M{"_id": validId}).
					Return(mockSingleResult)

				mockSingleResult.EXPECT().
					Decode(gomock.Any()).
					DoAndReturn(func(_ interface{}) error {
						return errors.New("failed to decode document")
					})
			},
			want: model.Product{},
			err:  error_wrapper.New(model.RErrMongoDBReadDocument, "failed to decode document"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF()
			result, err := repo.FindByID(tt.args.ctx, tt.args.productID)
			fmt.Println("INI RESULT DAN ERR", result, err)
			if err != nil {
				assert.Equal(t, model.Product{}, result)
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	type args struct {
		productID   string
		name        string
		ingredients []model.Ingredient
	}

	repo := product.NewProductRepository(mockClient, "testDB", "products")

	ingredients := []model.Ingredient{
		{
			ItemName: "Item A",
			ItemID:   "1231",
			Quantity: 10,
			Unit:     "kg",
		},
		{
			ItemID:   "1232",
			ItemName: "Item B",
			Quantity: 12,
			Unit:     "gr",
		},
	}

	tests := []struct {
		name  string
		mockF func(args) // Pass args explicitly
		args  args
		want  model.Product
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success Update a Document",
			args: args{
				name:        "Testing name",
				productID:   primitive.NewObjectID().Hex(),
				ingredients: ingredients,
			},
			mockF: func(a args) {
				objectID, _ := primitive.ObjectIDFromHex(a.productID) // Convert productID to ObjectID

				updatedData := bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "name", Value: a.name},
						{Key: "ingredients", Value: a.ingredients},
					}},
				}

				mockCollection.EXPECT().
					UpdateOne(gomock.Any(), bson.D{{Key: "_id", Value: objectID}}, updatedData).
					Return(&mongo.UpdateResult{MatchedCount: 1}, nil)
			},
			want: model.Product{
				Name:        "Testing name",
				Ingredients: ingredients,
			},
			err: nil,
		},
		{
			name: "Failed to convert productID to ObjectID",
			args: args{
				name:        "Invalid Product",
				productID:   "invalid_hex", // Invalid hex to trigger error
				ingredients: ingredients,
			},
			mockF: func(a args) {
			},
			want: model.Product{},
			err:  error_wrapper.New(model.RErrDecodeStringToObjectID, "the provided hex string is not a valid ObjectID"),
		},
		{
			name: "Failed to update document",
			args: args{
				name:        "New product",
				productID:   primitive.NewObjectID().Hex(),
				ingredients: ingredients,
			},
			mockF: func(a args) {
				objectID, _ := primitive.ObjectIDFromHex(a.productID) // Convert string to ObjectID

				updatedData := bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "name", Value: a.name},
						{Key: "ingredients", Value: a.ingredients},
					}},
				}

				mockCollection.EXPECT().
					UpdateOne(gomock.Any(), bson.D{{Key: "_id", Value: objectID}}, updatedData).
					Return(nil, errors.New("Failed to update a document")) // Simulate an update failure
			},
			err: error_wrapper.New(model.RErrMongoDBUpdateDocument, "Failed to update a document"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(tt.args) // Pass args to mockF
			err := repo.Update(context.Background(), tt.args.productID, tt.args.name, tt.args.ingredients)

			if err != nil {
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_mongodb.NewMockMongoDBClientWrapper(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabaseWrapper(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollectionWrapper(ctrl)

	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	repo := product.NewProductRepository(mockClient, "testDB", "products")

	type args struct {
		ctx       context.Context
		productID string
	}
	tests := []struct {
		name  string
		mockF func(args) // Pass args explicitly
		args  args
		err   *error_wrapper.ErrorWrapper
	}{
		{
			name: "Success delete",
			args: args{
				ctx:       context.Background(),
				productID: primitive.NewObjectID().Hex(),
			},
			mockF: func(a args) {
				objectId, _ := primitive.ObjectIDFromHex(a.productID)
				mockCollection.EXPECT().
					DeleteOne(a.ctx, bson.D{{Key: "_id", Value: objectId}}).
					Return(&mongo.DeleteResult{DeletedCount: 1}, nil)
			},
			err: nil,
		},
		{
			name: "Falied decode productID",
			args: args{
				ctx:       context.Background(),
				productID: "hehe",
			},
			mockF: func(a args) {},
			err:   error_wrapper.New(model.RErrDecodeStringToObjectID, "Error decoding product id"),
		},
		{
			name: "Failed deleting a document",
			args: args{
				ctx:       context.Background(),
				productID: primitive.NewObjectID().Hex(),
			},
			mockF: func(a args) {
				objectId, _ := primitive.ObjectIDFromHex(a.productID)
				mockCollection.EXPECT().
					DeleteOne(a.ctx, bson.D{{Key: "_id", Value: objectId}}).
					Return(nil, errors.New("Error deleting a document"))
			},
			err: error_wrapper.New(model.RErrMongoDBDeleteDocument, "Error deleting a document"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockF(tt.args)
			err := repo.Delete(tt.args.ctx, tt.args.productID)

			if err != nil {
				assert.Equal(t, tt.err.StatusCode(), err.StatusCode())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
