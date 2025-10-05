package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"financial-track/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation(getTimezone())
			return time.Now().In(loc)
		},
	})
	if err != nil {
		log.Fatal("❌ Error to connect to database: ", err)
	}
	DB = db

	tz := getTimezone()
	DB.Exec("SET TIME ZONE ?", tz)
	fmt.Println("✅ Database connected")
}

func getTimezone() string {
	tz := os.Getenv("APP_TIMEZONE")
	if tz == "" {
		tz = "America/Sao_Paulo"
	}
	return tz
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
