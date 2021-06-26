package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSecretKey() (secretKey string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return os.Getenv("SECRET_KEY")
}

func GetAdminKey() (secretKey string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return os.Getenv("ADMIN_KEY")
}