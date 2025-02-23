package itempurchasechain

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *itemPurchaseChainRepository) Create(ctx context.Context, itemID string, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper {
	itemPurchaseChain := model.ItemPurchaseChain{
		ItemID:       itemID,
		BranchID:     branchID,
		Purchase:     purchase,
		Quantity:     purchase.Quantity,
		Status:       model.StatusNotUsed,
		SalesRecords: []model.Sales{},
	}

	_, err := i.itemPurchaseChainCollection.InsertOne(
		ctx,
		itemPurchaseChain,
	)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return nil
}

// Find By Item Id dan BranchId
func (i *itemPurchaseChainRepository) Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	var (
		result []model.ItemPurchaseChainGet
	)

	filter := bson.M{}
	if payload.ItemID != "" {
		filter["item_id"] = payload.ItemID
	}

	if payload.BranchID != "" {
		filter["branch_id"] = payload.BranchID
	}

	if payload.Purchase.UUID != "" {
		filter["purchase.uuid"] = payload.Purchase.UUID
	}

	cur, err := i.itemPurchaseChainCollection.Find(ctx, filter)

	if err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	if err = cur.Decode(&result); err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	if len(result) == 0 {
		return nil, error_wrapper.New(model.RErrDataNotFound, "Item purchase chain not found")
	}

	return result, nil
}

// Update -> bisa update quantity / status

func (i *itemPurchaseChainRepository) Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper {
	filter := bson.D{{Key: "_id", Value: id}}

	updatedData := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "item_id", Value: payload.ItemID},
			{Key: "branch_id", Value: payload.BranchID},
			{Key: "purchase", Value: payload.Purchase},
			{Key: "quantity", Value: payload.Quantity},
			{Key: "status", Value: payload.Status},
			{Key: "sales", Value: payload.SalesRecords},
		}},
	}

	_, err := i.itemPurchaseChainCollection.UpdateOne(
		ctx,
		filter,
		updatedData,
	)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return nil
}
