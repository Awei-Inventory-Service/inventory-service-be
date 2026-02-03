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
	return &MongoCollection{collection: collection, database: r.database}
}

func (r *MongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, document, opts...)
}

func (r *MongoCollection) Database() MongoDBDatabaseWrapper {
	return &MongoDatabase{database: r.database}
}

func (d *MongoDatabase) StartSession(ctx context.Context) (MongoDBSessionWrapper, error) {
	session, err := d.database.Client().StartSession()
	if err != nil {
		return nil, err
	}
	return &MongoSession{session: session}, nil
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

func (r *MongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return r.collection.CountDocuments(ctx, filter, opts...)
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

func (s *MongoSession) StartTransaction() error {
	return s.session.StartTransaction()
}

// Commit the transaction
func (s *MongoSession) CommitTransaction(ctx context.Context) error {
	sc, ok := ctx.(mongo.SessionContext)
	if !ok {
		return mongo.ErrClientDisconnected
	}
	return sc.CommitTransaction(ctx)
}

// Abort (Rollback) the transaction
func (s *MongoSession) AbortTransaction(ctx context.Context) error {
	sc, ok := ctx.(mongo.SessionContext)
	if !ok {
		return mongo.ErrClientDisconnected
	}
	return sc.AbortTransaction(ctx)
}

// End the session
func (s *MongoSession) EndSession(ctx context.Context) {
	s.session.EndSession(ctx)
}

// Execute a function within a transaction
func (s *MongoSession) WithTransaction(ctx context.Context, fn func(sc mongo.SessionContext) error) error {
	return mongo.WithSession(ctx, s.session, func(sc mongo.SessionContext) error {
		if err := s.StartTransaction(); err != nil {
			return err
		}

		if err := fn(sc); err != nil {
			s.AbortTransaction(sc)
			return err
		}

		return s.CommitTransaction(sc)
	})
}
