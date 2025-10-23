package inventory_snapshot

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (i *inventorySnapshotResource) Create(ctx context.Context, payload model.InventorySnapshot) (errW *error_wrapper.ErrorWrapper) {
	_, err := i.inventorySnapshotCollection.InsertOne(ctx, payload)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}

	return
}

func (i *inventorySnapshotResource) Update(ctx context.Context, snapshotID string, payload model.InventorySnapshot) (errW *error_wrapper.ErrorWrapper) {
	id, err := primitive.ObjectIDFromHex(snapshotID)

	if err != nil {
		return error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: payload}}

	_, err = i.inventorySnapshotCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBUpdateDocument, err.Error())
	}

	return
}

func (i *inventorySnapshotResource) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.InventorySnapshot, *error_wrapper.ErrorWrapper) {
	var inventorySnapshots []model.InventorySnapshot

	// Build MongoDB filter
	mongoFilter := bson.D{}

	for _, f := range filter {
		if len(f.Values) == 0 {
			continue
		}

		switch f.Key {
		case "item_id":
			if len(f.Values) == 1 {
				mongoFilter = append(mongoFilter, bson.E{Key: "item_id", Value: f.Values[0]})
			} else {
				mongoFilter = append(mongoFilter, bson.E{Key: "item_id", Value: bson.D{{Key: "$in", Value: f.Values}}})
			}
		case "day":
			if len(f.Values) == 1 {
				dayInt := 0
				if day, err := strconv.Atoi(f.Values[0]); err == nil {
					dayInt = day
				}
				mongoFilter = append(mongoFilter, bson.E{Key: "day", Value: dayInt})
			}
		case "month":
			if len(f.Values) == 1 {
				monthInt := 0
				if month, err := strconv.Atoi(f.Values[0]); err == nil {
					monthInt = month
				}
				mongoFilter = append(mongoFilter, bson.E{Key: "month", Value: monthInt})
			}
		case "year":
			if len(f.Values) == 1 {
				yearInt := 0
				if year, err := strconv.Atoi(f.Values[0]); err == nil {
					yearInt = year
				}
				mongoFilter = append(mongoFilter, bson.E{Key: "year", Value: yearInt})
			}
		case "date":
			if len(f.Values) == 2 {
				// Date range filtering - expecting start_date and end_date
				startDate := f.Values[0]
				endDate := f.Values[1]
				mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{
					{Key: "$gte", Value: startDate},
					{Key: "$lte", Value: endDate},
				}})
			} else if len(f.Values) == 1 {
				// Single date filtering based on wildcard
				switch f.Wildcard {
				case "==":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: f.Values[0]})
				case ">=":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$gte", Value: f.Values[0]}}})
				case "<=":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$lte", Value: f.Values[0]}}})
				case ">":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$gt", Value: f.Values[0]}}})
				case "<":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$lt", Value: f.Values[0]}}})
				}
			}
		}
	}

	// Build sort options
	findOptions := options.Find()

	for _, o := range order {
		sortDirection := 1
		if !o.IsAsc {
			sortDirection = -1
		}
		findOptions.SetSort(bson.D{{Key: o.Key, Value: sortDirection}})
	}

	// Apply limit and offset
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}
	if offset > 0 {
		findOptions.SetSkip(int64(offset))
	}

	// Execute query
	cursor, err := i.inventorySnapshotCollection.Find(ctx, mongoFilter, findOptions)
	if err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	defer cursor.Close(ctx)

	// Decode results
	for cursor.Next(ctx) {
		var inventorySnapshot model.InventorySnapshot
		if err := cursor.Decode(&inventorySnapshot); err != nil {
			return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}
		inventorySnapshots = append(inventorySnapshots, inventorySnapshot)
	}

	if err := cursor.Err(); err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	return inventorySnapshots, nil
}

func (i *inventorySnapshotResource) FindByID(ctx context.Context, snapshotID string) (inventorySnapshot model.InventorySnapshot, errW *error_wrapper.ErrorWrapper) {
	id, err := primitive.ObjectIDFromHex(snapshotID)

	if err != nil {
		return model.InventorySnapshot{}, error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	result := i.inventorySnapshotCollection.FindOne(ctx, bson.M{
		"_id": id,
	})

	if err = result.Decode(&inventorySnapshot); err != nil {
		return model.InventorySnapshot{}, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	return
}

func (i *inventorySnapshotResource) Upsert(ctx context.Context, payload dto.CreateInventorySnapshotRequest) (errW *error_wrapper.ErrorWrapper) {
	today := time.Now().Truncate(24 * time.Hour)
	now := time.Now()

	filter := []dto.Filter{
		{
			Key:    "item_id",
			Values: []string{payload.ItemID},
		},
		{
			Key:    "branch_id",
			Values: []string{payload.BranchID},
		},
		{
			Key:    "day",
			Values: []string{fmt.Sprintf("%d", now.Day())},
		},
		{
			Key:    "month",
			Values: []string{fmt.Sprintf("%d", int(now.Month()))},
		},
		{
			Key:    "year",
			Values: []string{fmt.Sprintf("%d", now.Year())},
		},
	}

	existingInventorySnapshots, errW := i.Get(ctx, filter, []dto.Order{}, 1, 0)

	if errW != nil {
		return errW
	}

	// If snapshot exists, update it
	if len(existingInventorySnapshots) > 0 {
		existingSnapshot := existingInventorySnapshots[0]

		newValue := struct {
			Timestamp time.Time `json:"timestamp"`
			Value     float64   `json:"value"`
		}{
			Timestamp: time.Now(),
			Value:     payload.Value,
		}

		existingSnapshot.Values = append(existingSnapshot.Values, newValue)

		// Recalculate average
		var total float64
		for _, v := range existingSnapshot.Values {
			total += v.Value
		}
		existingSnapshot.Average = total / float64(len(existingSnapshot.Values))

		return i.Update(ctx, existingSnapshot.ID.Hex(), existingSnapshot)
	}
	// If snapshot doesn't exist, create new one
	newSnapshot := model.InventorySnapshot{
		ItemID:   payload.ItemID,
		BranchID: payload.BranchID,
		Date:     today,
		Average:  payload.Value,
		Day:      now.Day(),
		Month:    int(now.Month()),
		Year:     now.Year(),
		Values: []struct {
			Timestamp time.Time `json:"timestamp"`
			Value     float64   `json:"value"`
		}{
			{
				Timestamp: time.Now(),
				Value:     payload.Value,
			},
		},
	}

	return i.Create(ctx, newSnapshot)
}
