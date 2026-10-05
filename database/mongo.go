package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

func Connect() *mongo.Client {
	client, err := mongo.Connect(
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		fmt.Println("Ошибка подключения к MongoDB:", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		fmt.Println("MongoDB недоступна:", err)
	}

	fmt.Println("Подключение к MongoDB успешно")
	return client
}
