package handlers

import (
	"net/http"

	"github.com/TelvoWarrior/Sunday-Project/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPageData struct {
	Success bool
	Error   string
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	data := RegisterPageData{
		Success: r.URL.Query().Get("success") == "1",
		Error:   r.URL.Query().Get("error"),
	}

	renderTemplate(w, "templates/register.html", data)
}

func RegisterAccount(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		login := r.FormValue("login")
		password := r.FormValue("password")

		if login == "" || password == "" {
			http.Redirect(w, r, "/register?error=empty_fields", http.StatusSeeOther)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		collection := db.Collection("accounts")

		account := models.Account{
			Login:    login,
			Password: string(hashedPassword),
		}

		_, err = collection.InsertOne(r.Context(), account)

		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				http.Redirect(w, r, "/register?error=login_exists", http.StatusSeeOther)
				return
			}

			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/register?success=1", http.StatusSeeOther)
	}
}
