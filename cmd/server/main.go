package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TelvoWarrior/Sunday-Project/database"
	"github.com/TelvoWarrior/Sunday-Project/handlers"
)

func main() {
	fmt.Println("Hello, Sunday Project!")
	port := os.Getenv("PORT")

	if port == "" {
		port = "9999"
	}

	ctx := context.Background()

	client, err := database.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		databaseName = "sundayProject"
	}

	db := client.Database(databaseName)

	http.HandleFunc("GET /{$}", handlers.HomePage)
	http.HandleFunc("GET /login", handlers.LoginPage)
	http.HandleFunc("POST /login", handlers.LoginAccount(db))
	http.HandleFunc("GET /register", handlers.RegisterPage)
	http.HandleFunc("POST /register", handlers.RegisterAccount(db))
	http.HandleFunc("GET /profile", handlers.RequireAuth(db, handlers.Profile))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("HTTP server error:", err)
	}
}
