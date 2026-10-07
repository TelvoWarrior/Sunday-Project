package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TelvoWarrior/Sunday-Project/auth"
	"github.com/TelvoWarrior/Sunday-Project/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/login.html", nil)
}

func LoginAccount(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		login := r.FormValue("login")
		password := r.FormValue("password")

		if login == "" || password == "" {
			http.Redirect(w, r, "/login?error=empty_fields", http.StatusSeeOther)
			return
		}

		accountsCollection := db.Collection("accounts")

		var account models.Account

		err = accountsCollection.FindOne(
			r.Context(),
			bson.M{"login": login},
		).Decode(&account)

		fmt.Println("Account lookup error:", err)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
				return
			}

			http.Error(w, "Failed to sign in", http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(account.Password),
			[]byte(password),
		)

		if err != nil {
			http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
			return
		}

		token, err := auth.GenerateSessionToken()
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		sessionsCollection := db.Collection("sessions")

		session := models.Session{
			Token:     token,
			AccountID: account.ID,
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		}

		_, err = sessionsCollection.InsertOne(r.Context(), session)
		if err != nil {
			http.Error(w, "Failed to save session", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // For localhost only; For HTTPS-server — true.
			SameSite: http.SameSiteLaxMode,
			MaxAge:   7 * 24 * 60 * 60,
		})

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
