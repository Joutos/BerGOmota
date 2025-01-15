package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASS")
	dbname := os.Getenv("PG_DB_NAME")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

	}
	return db
}
