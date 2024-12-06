package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Миграция таблиц
	err = DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		return err
	}

	return nil
}
