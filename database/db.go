package database

import (
	"fmt"
	"log"
	"my-gram/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	godotenv.Load()
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbPort := os.Getenv("PGPORT")
	dbName := os.Getenv("PGDATABASE")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("Success connectiong to database")
	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Comment{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
