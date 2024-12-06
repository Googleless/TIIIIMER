package handlers

import (
	"html/template"
	"net/http"
)

// LoginHandler обрабатывает запросы на страницу входа
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		// Здесь будет обработка данных входа
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
