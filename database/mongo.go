package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

func Connect(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, fmt.Errorf("connect MongoDB: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, fmt.Errorf("connect MongoDB: %w", err)
	}

	fmt.Println("Подключение к MongoDB успешно")
	return client, nil
}
