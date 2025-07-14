package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads .env and .env_{ENV} files
func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ Could not load .env file")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	envFile := fmt.Sprintf(".env_%s", env)
	if err := godotenv.Overload(envFile); err != nil {
		log.Printf("⚠️ Could not load %s\n", envFile)
	}
}
