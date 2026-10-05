package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"sundayProject/auth"
	"sundayProject/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Println("Ошибка чтения шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Ошибка рендера шаблона:", err)
	}
}

func LoginAccount(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Не удалось прочитать форму", http.StatusBadRequest)
			return
		}

		login := r.FormValue("login")
		password := r.FormValue("password")

		if login == "" || password == "" {
			http.Redirect(w, r, "/login?error=empty_fields", http.StatusSeeOther)
			return
		}

		accountsCollection := client.Database("sundayProject").Collection("accounts")

		var account models.Account

		err = accountsCollection.FindOne(
			r.Context(),
			bson.M{"login": login},
		).Decode(&account)

		fmt.Println("Ошибка поиска аккаунта:", err)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				// Не сообщаем, существует ли такой логин.
				http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
				return
			}

			http.Error(w, "Не удалось выполнить вход", http.StatusInternalServerError)
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
			http.Error(w, "Не удалось создать сессию", http.StatusInternalServerError)
			return
		}

		sessionsCollection := client.Database("sundayProject").Collection("sessions")

		session := models.Session{
			Token:     token,
			AccountID: account.ID,
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		}

		_, err = sessionsCollection.InsertOne(r.Context(), session)
		if err != nil {
			http.Error(w, "Не удалось сохранить сессию", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // Для localhost; на HTTPS-сервере — true.
			SameSite: http.SameSiteLaxMode,
			MaxAge:   7 * 24 * 60 * 60,
		})

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
