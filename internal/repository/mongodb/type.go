package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDBCollection interface defines MongoDB collection operations
type MongoDBCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

// MongoDBDatabase interface defines database-level operations
type MongoDBDatabase interface {
	Collection(collectionName string, opts ...*options.CollectionOptions) MongoDBCollection
}

// MongoDBClient interface defines client-level operations
type MongoDBClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, readPref *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) MongoDBDatabase
}

// RealMongoDBClient wraps the actual MongoDB client
type RealMongoDBClient struct {
	client *mongo.Client
}

// RealMongoDBDatabase wraps the actual MongoDB database
type RealMongoDBDatabase struct {
	database *mongo.Database
}

type RealMongoDBCollection struct {
	collection *mongo.Collection
}

// NewRealMongoDBClient creates a new MongoDB client
