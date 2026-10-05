package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, filename string, data any) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Ошибка чтения шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Ошибка рендера шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}
