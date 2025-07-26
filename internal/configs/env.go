package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ Could not load .env file")
	}
}
