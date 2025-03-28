package itempurchasechain_test

import (
	"context"
	"testing"

	"github.com/inventory-service/app/model"
	itempurchasechain "github.com/inventory-service/app/usecase/item_purchase_chain"
	"github.com/inventory-service/lib/error_wrapper"
	mock_branch "github.com/inventory-service/mocks/resource/branch"
	mock_item "github.com/inventory-service/mocks/resource/item"
	mock_itempurchasechain "github.com/inventory-service/mocks/resource/item_purchase_chain"
	mock_purchase "github.com/inventory-service/mocks/resource/purchase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCalculateCost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	mockItemPurchaseChainRepo := mock_itempurchasechain.NewMockItemPurchaseChainResource(ctrl)
	mockItemRepo := mock_item.NewMockItemResource(ctrl)
	mockBranchRepo := mock_branch.NewMockBranchResource(ctrl)
	mockPurchaseRepo := mock_purchase.NewMockPurchaseResource(ctrl)

	itemPurchaseChainService := itempurchasechain.NewItemPurchaseChainService(
		mockItemPurchaseChainRepo,
		mockPurchaseRepo,
		mockItemRepo,
		mockBranchRepo,
	)
	itemId := "item-1"
	branchId := "branch-1"
	firstPurchaseChain := model.ItemPurchaseChainGet{
		ItemID:   itemId,
		BranchID: branchId,
		Quantity: 10,
		Status:   model.StatusInUse,
		Purchase: model.ItemPurchaseChainPurchase{
			PurchaseCost: 5.0,
		},
	}
	secondPurchaseChain := model.ItemPurchaseChainGet{
		ItemID:   itemId,
		BranchID: branchId,
		Quantity: 8,
		Status:   model.StatusNotUsed,
		Purchase: model.ItemPurchaseChainPurchase{
			PurchaseCost: 4.3,
		},
	}

	t.Run("Enough stock in first purchase chain", func(t *testing.T) {
		expectedItemPurchaseChain := []model.ItemPurchaseChainGet{
			{
				ID:       firstPurchaseChain.ID,
				ItemID:   firstPurchaseChain.ItemID,
				BranchID: firstPurchaseChain.BranchID,
				Quantity: 5,
				Status:   firstPurchaseChain.Status,
				Purchase: model.ItemPurchaseChainPurchase{
					PurchaseCost: 5.0,
				},
			},
		}
		mockItemPurchaseChainRepo.EXPECT().
			Get(ctx, gomock.Any()).
			Return([]model.ItemPurchaseChainGet{firstPurchaseChain}, nil)

		cost, purchaseChainResults, errW := itemPurchaseChainService.CalculateCost(ctx, itemId, branchId, 5)
		assert.Nil(t, errW)
		assert.Equal(t, expectedItemPurchaseChain, purchaseChainResults)
		assert.Equal(t, 25.0, cost)
	})

	t.Run("Not enough stock in second purchase chain", func(t *testing.T) {
		expectedItemPurchaseChain := []model.ItemPurchaseChainGet{
			{
				ID:       firstPurchaseChain.ID,
				ItemID:   firstPurchaseChain.ItemID,
				BranchID: firstPurchaseChain.BranchID,
				Quantity: 0,
				Status:   model.StatusUsed,
				Purchase: firstPurchaseChain.Purchase,
			},
			{
				ID:       secondPurchaseChain.ID,
				ItemID:   secondPurchaseChain.ItemID,
				BranchID: secondPurchaseChain.BranchID,
				Quantity: 6,
				Status:   model.StatusInUse,
				Purchase: secondPurchaseChain.Purchase,
			},
		}
		gomock.InOrder(
			mockItemPurchaseChainRepo.EXPECT().Get(ctx, gomock.Any()).Return([]model.ItemPurchaseChainGet{firstPurchaseChain}, nil).Times(1),
			mockItemPurchaseChainRepo.EXPECT().Get(ctx, gomock.Any()).Return([]model.ItemPurchaseChainGet{secondPurchaseChain}, nil).Times(1),
		)
		cost, purchaseChainResults, errW := itemPurchaseChainService.CalculateCost(ctx, itemId, branchId, 12)
		assert.Nil(t, errW)
		assert.Equal(t, expectedItemPurchaseChain, purchaseChainResults)
		assert.Equal(t, 58.6, cost)
	})

	t.Run("No purchase chain with status in use", func(t *testing.T) {
		itemPurchaseChain := []model.ItemPurchaseChainGet{
			{
				ID:       "purchase-1",
				ItemID:   itemId,
				BranchID: branchId,
				Quantity: 10,
				Status:   model.StatusNotUsed,
				Purchase: model.ItemPurchaseChainPurchase{
					PurchaseCost: 5.0,
				},
			},
		}
		gomock.InOrder(
			mockItemPurchaseChainRepo.EXPECT().Get(ctx, gomock.Any()).Return(nil, error_wrapper.New(model.RErrDataNotFound, gomock.Any())),
			mockItemPurchaseChainRepo.EXPECT().Get(ctx, gomock.Any()).Return(itemPurchaseChain, nil),
			mockItemPurchaseChainRepo.EXPECT().
				Update(ctx, "purchase-1", gomock.Any()).
				Do(func(_ context.Context, id string, item model.ItemPurchaseChain) {
					assert.Equal(t, "purchase-1", id)
					assert.Equal(t, model.StatusInUse, item.Status)
					assert.Equal(t, itemPurchaseChain[0].ItemID, item.ItemID)
				}).
				Return(nil),
		)
		cost, _, errW := itemPurchaseChainService.CalculateCost(ctx, itemId, branchId, 5)
		assert.Nil(t, errW)
		assert.Equal(t, 25.0, cost)
	})
}
