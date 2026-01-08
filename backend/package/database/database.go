package database

import (
	"ezkost/internal/config"
	"ezkost/internal/repository/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Room{},
		&model.Tenant{},
		&model.Payment{},
		&model.Expense{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
}
