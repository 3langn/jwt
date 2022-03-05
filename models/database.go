package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)


var db *gorm.DB

func GetDb() *gorm.DB{
	return db
}

func Setup() {
	var err error

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Fail to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

