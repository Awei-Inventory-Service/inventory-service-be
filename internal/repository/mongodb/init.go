package mongodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewRealMongoDBClient(client *mongo.Client) MongoDBClient {
	return &RealMongoDBClient{client: client}
}

func InitMongoDB() (MongoDBClient, error) {
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoDatabaseName := os.Getenv("MONGO_DATABASE")

	if mongoUser == "" || mongoPassword == "" || mongoHost == "" || mongoPort == "" {
		return nil, errors.New("username, password, host, or port can't be empty")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", mongoUser, mongoPassword, mongoHost, mongoPort, mongoDatabaseName)
	clientOpts := options.Client().ApplyURI(uri)

	mongoClient, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return NewRealMongoDBClient(mongoClient), nil
}
