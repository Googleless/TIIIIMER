package main

import (
	"TIIIIMER/handlers"
	"TIIIIMER/models"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	// Регистрация маршрутов
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	// Запуск веб-сервера
	log.Println("Сервер запущен на http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
