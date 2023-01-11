package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%v TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("TZ"))

}
func Connect() {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
