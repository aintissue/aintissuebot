package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

func Init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

func Run() {
	log.Println(os.Getenv("TELEGRAM_KEY"))
}
