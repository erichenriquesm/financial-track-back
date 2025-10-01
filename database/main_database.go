package database

import (
	"fmt"
	"log"
	"os"

	"financial-track/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error to connect to database: ", err)
	}
	DB = db
	fmt.Println("✅ Database connected")
}

func Migrate() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Expense{},
	)

	if err != nil {
		log.Fatal("❌ Error to run migrations: ", err)
	}

	fmt.Println("📦 Migrations applied")
}
