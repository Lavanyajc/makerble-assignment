package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
        
)

var DB *sql.DB

func ConnectDB() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Get DB details from environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// PostgreSQL connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	// Open the DB connection
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("❌ Unable to open DB connection: ", err)
	}

	// Ping DB to test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Cannot connect to DB: ", err)
	}

	fmt.Println("✅Hii JC , Successfully connected to PostgreSQL !!")
}

