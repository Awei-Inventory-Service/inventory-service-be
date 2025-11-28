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
				startDate, err1 := time.Parse("2006-01-02 15:04:05", f.Values[0])
				endDate, err2 := time.Parse("2006-01-02 15:04:05", f.Values[1])
				if err1 == nil && err2 == nil {
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{
						{Key: "$gte", Value: startDate},
						{Key: "$lte", Value: endDate},
					}})
				}
			} else if len(f.Values) == 1 {
				// Parse the date string to time.Time - try multiple formats
				var dateValue time.Time
				var err error

				// Try with time format first
				dateValue, err = time.Parse("2006-01-02 15:04:05", f.Values[0])
				if err != nil {
					// Try date-only format
					dateValue, err = time.Parse("2006-01-02", f.Values[0])
					if err != nil {
						fmt.Println("Error parsing time in get inventory snapshot resource", f.Values[0])
						continue
					}
				}

				// Single date filtering based on wildcard
				switch f.Wildcard {
				case "==":
					// For equality, use date range for the entire day
					startOfDay := time.Date(dateValue.Year(), dateValue.Month(), dateValue.Day(), 0, 0, 0, 0, dateValue.Location())
					endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{
						{Key: "$gte", Value: startOfDay},
						{Key: "$lte", Value: endOfDay},
					}})
				case ">=":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$gte", Value: dateValue}}})
				case "<=":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$lte", Value: dateValue}}})
				case ">":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$gt", Value: dateValue}}})
				case "<":
					mongoFilter = append(mongoFilter, bson.E{Key: "date", Value: bson.D{{Key: "$lt", Value: dateValue}}})
				}
			}
		}
	}

	fmt.Println("Mongodb filter", mongoFilter)
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
			Values: []string{fmt.Sprintf("%d", payload.Date.Day())},
		},
		{
			Key:    "month",
			Values: []string{fmt.Sprintf("%d", int(payload.Date.Month()))},
		},
		{
			Key:    "year",
			Values: []string{fmt.Sprintf("%d", payload.Date.Year())},
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
		existingSnapshot.Latest = payload.Value
		existingSnapshot.Balance = payload.Balance
		return i.Update(ctx, existingSnapshot.ID.Hex(), existingSnapshot)
	}
	// If snapshot doesn't exist, create new one
	newSnapshot := model.InventorySnapshot{
		ItemID:   payload.ItemID,
		BranchID: payload.BranchID,
		Date:     payload.Date,
		Balance:  payload.Balance,
		Average:  payload.Value,
		Latest:   payload.Value,
		Day:      payload.Date.Day(),
		Month:    int(payload.Date.Month()),
		Year:     payload.Date.Year(),
		Values: []struct {
			Timestamp time.Time `json:"timestamp"`
			Value     float64   `json:"value"`
		}{
			{
				Timestamp: payload.Date,
				Value:     payload.Value,
			},
		},
	}

	return i.Create(ctx, newSnapshot)
}

func (i *inventorySnapshotResource) GetPreviousDaySnapshot(ctx context.Context, targetTime time.Time, branchID, itemID string) (*model.InventorySnapshot, *error_wrapper.ErrorWrapper) {
	previousDay := targetTime.AddDate(0, 0, -1)

	filter := []dto.Filter{
		{
			Key:    "item_id",
			Values: []string{itemID},
		},
		{
			Key:    "branch_id",
			Values: []string{branchID},
		},
		{
			Key:    "day",
			Values: []string{fmt.Sprintf("%d", previousDay.Day())},
		},
		{
			Key:    "month",
			Values: []string{fmt.Sprintf("%d", int(previousDay.Month()))},
		},
		{
			Key:    "year",
			Values: []string{fmt.Sprintf("%d", previousDay.Year())},
		},
	}

	snapshots, errW := i.Get(ctx, filter, []dto.Order{}, 1, 0)
	if errW != nil {
		return nil, errW
	}

	if len(snapshots) == 0 {
		return nil, error_wrapper.New(model.RErrDataNotFound, "no snapshot found for previous day")
	}

	return &snapshots[0], nil
}

func (i *inventorySnapshotResource) GetSnapshotBasedOndDate(ctx context.Context, date time.Time) (model.InventorySnapshot, *error_wrapper.ErrorWrapper) {
	filter := []dto.Filter{
		{
			Key:      "day",
			Values:   []string{fmt.Sprintf("%d", date.Day())},
			Wildcard: "==",
		},
		{
			Key:      "month",
			Values:   []string{fmt.Sprintf("%d", int(date.Month()))},
			Wildcard: "==",
		},
		{
			Key:    "year",
			Values: []string{fmt.Sprintf("%d", date.Year())},
		},
	}

	snapshot, errW := i.Get(ctx, filter, []dto.Order{}, 0, 0)
	if errW != nil {
		fmt.Println("Error getting snapshot based on date", errW)
		return model.InventorySnapshot{}, errW
	}

	if len(snapshot) == 0 {
		errW = error_wrapper.New(model.RErrDataNotFound, "Inventory snapshot not found")
		return model.InventorySnapshot{}, errW
	}

	return snapshot[0], nil
}
