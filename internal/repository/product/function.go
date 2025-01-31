package product

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *productRepository) Create(ctx context.Context, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper {
	product := model.Product{
		Name:        name,
		Ingredients: ingredients,
	}

	_, err := p.productCollection.InsertOne(ctx, product)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBCreateDocument, err.Error())
	}
	return nil
}

func (p *productRepository) FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper) {

	cursor, err := p.productCollection.Find(ctx, bson.M{})
	if err != nil {
		errW := error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		return nil, errW
	}

	defer cursor.Close(ctx)

	var products []model.Product
	for cursor.Next(ctx) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	return products, nil
}

func (p *productRepository) FindByID(ctx context.Context, productID string) (model.Product, *error_wrapper.ErrorWrapper) {

	objectID, err := primitive.ObjectIDFromHex(productID)

	if err != nil {
		return model.Product{}, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}

	result := p.productCollection.FindOne(ctx, bson.M{
		"_id": objectID,
	})

	var product model.Product

	if err = result.Decode(&product); err != nil {
		return model.Product{}, error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
	}
	return product, nil
}

func (p *productRepository) Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper {
	id, err := primitive.ObjectIDFromHex(productID)

	if err != nil {
		return error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}

	updatedData := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "ingredients", Value: ingredients},
		}},
	}
	_, err = p.productCollection.UpdateOne(ctx, filter, updatedData)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBUpdateDocument, err.Error())
	}

	return nil
}

func (p *productRepository) Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper {

	id, err := primitive.ObjectIDFromHex(productID)

	if err != nil {
		return error_wrapper.New(model.RErrDecodeStringToObjectID, err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}

	_, err = p.productCollection.DeleteOne(ctx, filter)

	if err != nil {
		return error_wrapper.New(model.RErrMongoDBDeleteDocument, err.Error())
	}

	return nil
}
