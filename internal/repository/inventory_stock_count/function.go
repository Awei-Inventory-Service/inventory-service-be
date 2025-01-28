package inventorystockcount

import (
	"context"
	"time"

	"github.com/inventory-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *inventoryStockCountRepository) Create(ctx context.Context, branchID string, items []model.ItemCount) error {
	newData := model.InventoryStockCount{
		BranchID:  branchID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Items:     items,
	}

	_, err := i.inventoryStockCountCollection.InsertOne(ctx, newData)

	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryStockCountRepository) Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) error {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return err
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
		return err
	}

	return nil
}

func (i *inventoryStockCountRepository) FindAll(ctx context.Context) ([]model.InventoryStockCount, error) {
	cur, err := i.inventoryStockCountCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var inventoryStockCounts []model.InventoryStockCount
	for cur.Next(ctx) {
		var inventoryStockCount model.InventoryStockCount

		if err := cur.Decode(&inventoryStockCount); err != nil {
			return nil, err
		}
		inventoryStockCounts = append(inventoryStockCounts, inventoryStockCount)
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountRepository) FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, error) {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return model.InventoryStockCount{}, err
	}

	result := i.inventoryStockCountCollection.FindOne(ctx, bson.M{
		"_id": id,
	})

	var inventoryStockCount model.InventoryStockCount

	if err = result.Decode(&inventoryStockCount); err != nil {
		return model.InventoryStockCount{}, err
	}

	return inventoryStockCount, err
}

func (i *inventoryStockCountRepository) FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, error) {
	cursor, err := i.inventoryStockCountCollection.Find(ctx, bson.D{
		{Key: "branch_id", Value: branchID},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var inventoryStockCounts []model.InventoryStockCount

	for cursor.Next(ctx) {
		var inventoryStockCount model.InventoryStockCount

		if err := cursor.Decode(&inventoryStockCount); err != nil {
			return nil, err
		}

		inventoryStockCounts = append(inventoryStockCounts, inventoryStockCount)
	}

	return inventoryStockCounts, nil
}

func (i *inventoryStockCountRepository) Delete(ctx context.Context, stockCountID string) error {
	id, err := primitive.ObjectIDFromHex(stockCountID)

	if err != nil {
		return err
	}

	_, err = i.inventoryStockCountCollection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: id},
	})

	if err != nil {
		return err
	}

	return nil
}
