package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TelvoWarrior/Sunday-Project/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Вы успешно вошли в аккаунт")
}

func RequireAuth(db *mongo.Database, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		sessionsCollection := db.Collection("sessions")

		var session models.Session

		err = sessionsCollection.FindOne(
			r.Context(),
			bson.M{
				"token":      cookie.Value,
				"expires_at": bson.M{"$gt": time.Now()},
			},
		).Decode(&session)

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
