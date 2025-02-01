package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (r *MongoSingleResult) Decode(v interface{}) error {
	return r.singleResult.Decode(v)
}

func (m *MongoClient) Connect(ctx context.Context) error {
	return m.client.Connect(ctx)
}

func (m *MongoClient) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoClient) Ping(ctx context.Context, readPref *readpref.ReadPref) error {
	return m.client.Ping(ctx, readPref)
}

func (r *MongoClient) Database(name string, opts ...*options.DatabaseOptions) MongoDBDatabaseWrapper {
	database := r.client.Database(name, opts...)
	return &MongoDatabase{database: database}
}

func (r *MongoDatabase) Collection(collectionName string, opts ...*options.CollectionOptions) MongoDBCollectionWrapper {
	collection := r.database.Collection(collectionName, opts...)
	return &MongoCollection{collection: collection}
}

func (r *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, document, opts...)
}

func (r *MongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoDBCursorWrapper, error) {
	cursor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &MongoCursor{cursor: cursor}, nil
}

func (r *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, filter, update, opts...)
}

func (r *MongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResultWrapper {
	return &MongoSingleResult{singleResult: r.collection.FindOne(ctx, filter, opts...)}
}

func (r *MongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, filter, opts...)
}

func (c *MongoCursor) Next(ctx context.Context) bool {
	return c.cursor.Next(ctx)
}

func (c *MongoCursor) Decode(val interface{}) error {
	return c.cursor.Decode(val)
}

func (c *MongoCursor) Err() error {
	return c.cursor.Err()
}

func (c *MongoCursor) Close(ctx context.Context) error {
	return c.cursor.Close(ctx)
}
