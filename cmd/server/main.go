package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TelvoWarrior/Sunday-Project/database"
	"github.com/TelvoWarrior/Sunday-Project/handlers"
)

const port string = ":9999"

func main() {
	fmt.Println("Hello, Sunday project!")
	client := database.Connect()

	defer client.Disconnect(context.Background())

	http.HandleFunc("GET /{$}", handlers.HomePage)
	http.HandleFunc("GET /login", handlers.LoginPage)
	http.HandleFunc("POST /login", handlers.LoginAccount(client))
	http.HandleFunc("GET /register", handlers.RegisterPage)
	http.HandleFunc("POST /register", handlers.RegisterAccount(client))
	http.HandleFunc("GET /profile", handlers.RequireAuth(client, handlers.Profile))

	_ = http.ListenAndServe(port, nil)
}
