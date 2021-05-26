package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSecretKey() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	return secretKey
}