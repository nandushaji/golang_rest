package config

import (
	"fmt"

	"github.com/nandushaji/golang_rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=golang_db password=qwerty dbname=restdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	fmt.Printf("Connected to %v", db)
	return db, err
}
