package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, filename string, data any) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Failed to load template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
