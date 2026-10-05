package handlers

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/home.html", nil)
}
