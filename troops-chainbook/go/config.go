package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value from key ---
func Config(key string) string {

	// load database config file
	err := godotenv.Load("env")

	if err != nil {
		log.Println("Error loading env file")
	}

	return os.Getenv(key)
}
