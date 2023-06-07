package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	fmt.Println("Loading env!!")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
	fmt.Println("Success!!")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(user, password)
}
