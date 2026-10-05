package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"sundayProject/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPageData struct {
	Success bool
	Error   string
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		fmt.Println("Ошибка чтения шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := RegisterPageData{
		Success: r.URL.Query().Get("success") == "1",
		Error:   r.URL.Query().Get("error"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Ошибка рендера шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}

func RegisterAccount(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Не удалось прочитать форму", http.StatusBadRequest)
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
			http.Error(w, "Не удалось обработать пароль", http.StatusInternalServerError)
			return
		}

		collection := client.Database("sundayProject").Collection("accounts")

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

			http.Error(w, "Не удалось создать пользователя", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/register?success=1", http.StatusSeeOther)
	}
}
