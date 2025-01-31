package product_test

import (
	"testing"

	"github.com/inventory-service/internal/repository/product"
	mock_mongodb "github.com/inventory-service/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks using NewMockXXX methods
	mockClient := mock_mongodb.NewMockMongoDBClient(ctrl)
	mockDatabase := mock_mongodb.NewMockMongoDBDatabase(ctrl)
	mockCollection := mock_mongodb.NewMockMongoDBCollection(ctrl)

	// Set up expectations
	mockClient.EXPECT().
		Database("testDB", gomock.Any()).
		Return(mockDatabase)

	mockDatabase.EXPECT().
		Collection("products", gomock.Any()).
		Return(mockCollection)

	// Create repository
	repo := product.NewProductRepository(mockClient, "testDB", "products")

	// Assert repo is not nil
	assert.NotNil(t, repo)
}
