package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (m *RealMongoDBClient) Connect(ctx context.Context) error {
	return m.client.Connect(ctx)
}

func (m *RealMongoDBClient) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *RealMongoDBClient) Ping(ctx context.Context, readPref *readpref.ReadPref) error {
	return m.client.Ping(ctx, readPref)
}

func (r *RealMongoDBClient) Database(name string, opts ...*options.DatabaseOptions) MongoDBDatabase {
	database := r.client.Database(name, opts...)
	return &RealMongoDBDatabase{database: database}
}

func (r *RealMongoDBDatabase) Collection(collectionName string, opts ...*options.CollectionOptions) MongoDBCollection {
	collection := r.database.Collection(collectionName, opts...)
	return &RealMongoDBCollection{collection: collection}
}

func (r *RealMongoDBCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, document, opts...)
}

func (r *RealMongoDBCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	return r.collection.Find(ctx, filter, opts...)
}

func (r *RealMongoDBCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, filter, update, opts...)
}

func (r *RealMongoDBCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return r.collection.FindOne(ctx, filter, opts...)
}

func (r *RealMongoDBCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, filter, opts...)
}
