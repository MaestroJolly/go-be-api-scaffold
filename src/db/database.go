package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MaestroJolly/go-be-api-scaffold/src/db/constants"
)

var Database *gorm.DB

// function to connect to the database
func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	timezone := os.Getenv("DB_TIMEZONE")

	env := os.Getenv("ENV")

	// connect to remote staging/dev database if env is set to development
	if env == constants.Development {
		host = os.Getenv("DEV_DB_HOST")
		username = os.Getenv("DEV_DB_USER")
		password = os.Getenv("DEV_DB_PASSWORD")
		databaseName = os.Getenv("DEV_DB_NAME")
		port = os.Getenv("DEV_DB_PORT")
		timezone = os.Getenv("DEV_DB_TIMEZONE")
	}

	// connect to remote staging/dev database if env is set to production
	if env == constants.Production {
		host = os.Getenv("PROD_DB_HOST")
		username = os.Getenv("PROD_DB_USER")
		password = os.Getenv("PROD_DB_PASSWORD")
		databaseName = os.Getenv("PROD_DB_NAME")
		port = os.Getenv("PROD_DB_PORT")
		timezone = os.Getenv("PROD_DB_TIMEZONE")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, username, password, databaseName, port, timezone)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Error connecting to database, %v", err))
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
