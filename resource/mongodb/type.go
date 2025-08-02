package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBSingleResultWrapper interface {
	Decode(v interface{}) error
}

type MongoDBCursorWrapper interface {
	Next(ctx context.Context) bool
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
}

type MongoDBCollectionWrapper interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoDBCursorWrapper, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResultWrapper
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Database() MongoDBDatabaseWrapper // Add this method
}

type MongoDBSessionWrapper interface {
	StartTransaction() error
	CommitTransaction(ctx context.Context) error
	AbortTransaction(ctx context.Context) error
	EndSession(ctx context.Context)
	WithTransaction(ctx context.Context, fn func(sc mongo.SessionContext) error) error
}

type MongoSession struct {
	session mongo.Session
}

type MongoDBDatabaseWrapper interface {
	Collection(collectionName string, opts ...*options.CollectionOptions) MongoDBCollectionWrapper
	StartSession(ctx context.Context) (MongoDBSessionWrapper, error)
}

type MongoDBClientWrapper interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, readPref *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) MongoDBDatabaseWrapper
}

type MongoClient struct {
	client *mongo.Client
}

type MongoDatabase struct {
	database *mongo.Database
}

type MongoCollection struct {
	collection *mongo.Collection
	database   *mongo.Database
}

type MongoCursor struct {
	cursor *mongo.Cursor
}

type MongoSingleResult struct {
	singleResult *mongo.SingleResult
}
