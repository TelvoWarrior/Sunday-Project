package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, fmt.Errorf("connect MongoDB: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, fmt.Errorf("connect MongoDB: %w", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
