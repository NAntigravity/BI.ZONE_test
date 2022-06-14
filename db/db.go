package db

import (
	"BI.ZONE_test/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func init() {
	_ = godotenv.Load()
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	var err error
	db, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatalf("Failed to open connection to database. %v", err)
	}
	_ = db.AutoMigrate(&models.User{}, &models.Event{})
	_ = db
}
