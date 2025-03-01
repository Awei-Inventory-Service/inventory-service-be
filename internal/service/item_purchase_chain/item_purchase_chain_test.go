package itempurchasechain_test

import (
	"context"
	"testing"

	"github.com/inventory-service/internal/model"
	itempurchasechain "github.com/inventory-service/internal/service/item_purchase_chain"
	mock_branch "github.com/inventory-service/mocks/repository/branch"
	mock_item "github.com/inventory-service/mocks/repository/item"
	mock_itempurchasechain "github.com/inventory-service/mocks/repository/item_purchase_chain"
	mock_purchase "github.com/inventory-service/mocks/repository/purchase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCalculateCost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	mockItemPurchaseChainRepo := mock_itempurchasechain.NewMockItemPurchaseChainRepository(ctrl)
	mockItemRepo := mock_item.NewMockItemRepository(ctrl)
	mockBranchRepo := mock_branch.NewMockBranchRepository(ctrl)
	mockPurchaseRepo := mock_purchase.NewMockPurchaseRepository(ctrl)

	itemPurchaseChainService := itempurchasechain.NewItemPurchaseChainService(
		mockItemPurchaseChainRepo,
		mockPurchaseRepo,
		mockItemRepo,
		mockBranchRepo,
	)
	itemId := "item-1"
	branchId := "branch-1"
	firstPurchaseChain := []model.ItemPurchaseChainGet{
		{
			ItemID:   itemId,
			BranchID: branchId,
			Quantity: 10,
			Status:   model.StatusInUse,
			Purchase: model.Purchase{
				Item: model.Item{
					Price: 5.0,
				},
			},
		},
	}

	t.Run("Enough stock in first purchase chain", func(t *testing.T) {
		expectedItemPurchaseChain := []model.ItemPurchaseChainGet{
			{
				ID:       firstPurchaseChain[0].ID,
				ItemID:   firstPurchaseChain[0].ItemID,
				BranchID: firstPurchaseChain[0].BranchID,
				Quantity: 5,
				Status:   firstPurchaseChain[0].Status,
				Purchase: model.Purchase{
					Item: model.Item{
						Price: 5.0,
					},
				},
			},
		}
		mockItemPurchaseChainRepo.EXPECT().
			Get(ctx, gomock.Any()).
			Return(firstPurchaseChain, nil)

		cost, purchaseChainResults, errW := itemPurchaseChainService.CalculateCost(ctx, itemId, branchId, 5)
		assert.Nil(t, errW)
		assert.Equal(t, expectedItemPurchaseChain, purchaseChainResults)
		assert.Equal(t, 25.0, cost)
	})
}
