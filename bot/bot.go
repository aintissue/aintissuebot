package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Package init function
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

// Prepares the environment and runs the bot
func Run() {
	log.Println(os.Getenv("TELEGRAM_KEY"))
}
