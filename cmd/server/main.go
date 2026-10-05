package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/TelvoWarrior/Sunday-Project/database"
	"github.com/TelvoWarrior/Sunday-Project/handlers"
)

const port string = ":9999"

func main() {
	fmt.Println("Hello, Sunday project!")

	ctx := context.Background()

	client, err := database.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	http.HandleFunc("GET /{$}", handlers.HomePage)
	http.HandleFunc("GET /login", handlers.LoginPage)
	http.HandleFunc("POST /login", handlers.LoginAccount(client))
	http.HandleFunc("GET /register", handlers.RegisterPage)
	http.HandleFunc("POST /register", handlers.RegisterAccount(client))
	http.HandleFunc("GET /profile", handlers.RequireAuth(client, handlers.Profile))

	_ = http.ListenAndServe(port, nil)
}
