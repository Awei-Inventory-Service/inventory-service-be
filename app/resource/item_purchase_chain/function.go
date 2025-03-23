package itempurchasechain

import (
	"context"
	"fmt"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *itemPurchaseChainResource) Create(ctx context.Context, itemPurchaseChain model.ItemPurchaseChain) *error_wrapper.ErrorWrapper {
	_, err := i.itemPurchaseChainCollection.InsertOne(
		ctx,
		itemPurchaseChain,
	)

	if err != nil {
		fmt.Println("Error", err.Error())
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return nil
}

// Find By Item Id dan BranchId
func (i *itemPurchaseChainResource) Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
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

	for cur.Next(ctx) {
		var itemPurchaseChain model.ItemPurchaseChain

		if err = cur.Decode(&itemPurchaseChain); err != nil {
			return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}
		result = append(result, model.ItemPurchaseChainGet{
			ID:       itemPurchaseChain.UUID.Hex(),
			ItemID:   itemPurchaseChain.ItemID,
			BranchID: itemPurchaseChain.BranchID,
			Purchase: itemPurchaseChain.Purchase,
			Quantity: itemPurchaseChain.Quantity,
			Status:   itemPurchaseChain.Status,
			Sales:    itemPurchaseChain.Sales,
		})
	}

	if len(result) == 0 {
		return nil, error_wrapper.New(model.RErrDataNotFound, "Item purchase chain not found")
	}

	return result, nil
}

// Update -> bisa update quantity / status

func (i *itemPurchaseChainResource) BulkUpdate(ctx context.Context, payload []model.ItemPurchaseChainGet) *error_wrapper.ErrorWrapper {
	session, err := i.itemPurchaseChainCollection.Database().StartSession(ctx)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCollection, err.Error())
	}

	err = session.WithTransaction(ctx, func(sc mongo.SessionContext) error {
		for _, item := range payload {
			filter := bson.D{{Key: "_id", Value: item.ID}}

			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "item_id", Value: item.ItemID},
					{Key: "branch_id", Value: item.BranchID},
					{Key: "purchase", Value: item.Purchase},
					{Key: "quantity", Value: item.Quantity},
					{Key: "status", Value: item.Status},
					{Key: "sales", Value: item.Sales},
				}},
			}

			_, err := i.itemPurchaseChainCollection.UpdateOne(sc, filter, update)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBUpdateDocument, err.Error())
	}
	return nil
}

func (i *itemPurchaseChainResource) Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "item_id", Value: payload.ItemID},
				{Key: "branch_id", Value: payload.BranchID},
				{Key: "purchase", Value: payload.Purchase},
				{Key: "quantity", Value: payload.Quantity},
				{Key: "status", Value: payload.Status},
				{Key: "sales", Value: payload.Sales},
			}},
		}},
	}
	_, err := i.itemPurchaseChainCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBUpdateDocument, err.Error())
	}

	return nil
}
