package inventorystockcount

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *inventoryStockCountResource) Create(ctx context.Context, newData model.InventoryStockCount) *error_wrapper.ErrorWrapper {
	_, err := i.inventoryStockCountCollection.InsertOne(ctx, newData)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return nil
}

func (i *inventoryStockCountResource) Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}

	updatedData := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "branch_id", Value: branchID},
			{Key: "items", Value: items},
		}},
	}

	_, err = i.inventoryStockCountCollection.UpdateOne(ctx, filter, updatedData)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBUpdateDocument, err.Error())
	}

	return nil
}

func (i *inventoryStockCountResource) FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	cur, err := i.inventoryStockCountCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	defer cur.Close(ctx)

	var inventoryStockCounts []model.InventoryStockCount
	for cur.Next(ctx) {
		var inventoryStockCount model.InventoryStockCount

		if err := cur.Decode(&inventoryStockCount); err != nil {
			return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}
		inventoryStockCounts = append(inventoryStockCounts, inventoryStockCount)
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountResource) FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return model.InventoryStockCount{}, error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	result := i.inventoryStockCountCollection.FindOne(ctx, bson.M{
		"_id": id,
	})

	var inventoryStockCount model.InventoryStockCount

	if err = result.Decode(&inventoryStockCount); err != nil {
		return model.InventoryStockCount{}, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	return inventoryStockCount, nil
}

func (i *inventoryStockCountResource) FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper) {
	cursor, err := i.inventoryStockCountCollection.Find(ctx, bson.D{
		{Key: "branch_id", Value: branchID},
	})

	if err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	defer cursor.Close(ctx)

	var inventoryStockCounts []model.InventoryStockCount

	for cursor.Next(ctx) {
		var inventoryStockCount model.InventoryStockCount

		if err := cursor.Decode(&inventoryStockCount); err != nil {
			return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}

		inventoryStockCounts = append(inventoryStockCounts, inventoryStockCount)
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountResource) Delete(ctx context.Context, stockCountID string) *error_wrapper.ErrorWrapper {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	_, err = i.inventoryStockCountCollection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: id},
	})

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBDeleteDocument, err.Error())
	}

	return nil
}
