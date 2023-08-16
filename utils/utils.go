package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetValue returns the configuration value for the given key from the environment variables
func GetValue(key string) string {
	// Load the .env file
	err := godotenv.Load(".env")

	// If error occurred while loading the .env file
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	// Return the value for the given key
	return os.Getenv(key)
}
