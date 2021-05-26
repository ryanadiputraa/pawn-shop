package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func OpenConnection() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB_NAME")

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := sql.Open("postgres", desc)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} 

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}