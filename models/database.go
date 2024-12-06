package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB // Переменная для хранения подключения к базе данных

// InitDatabase инициализирует базу данных и выполняет миграции
func InitDatabase() error {
	var err error

	// Подключение к базе данных SQLite
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Миграция схемы для таблиц
	err = DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		return err
	}

	return nil
}

// Модель пользователя
type User struct {
	ID       uint   `gorm:"primaryKey"`      // Первичный ключ
	Username string `gorm:"unique;not null"` // Уникальное имя пользователя
	Password string `gorm:"not null"`        // Пароль
	Role     string `gorm:"not null"`        // Роль (admin или user)
}

// Модель задачи
type Task struct {
	ID          uint   `gorm:"primaryKey"` // Первичный ключ
	Description string `gorm:"not null"`   // Описание задачи
	Completed   bool   // Статус выполнения
	UserID      uint   // Связь с пользователем
}
