package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for start bot command
func startCommand(c telebot.Context) error {
	log.Println(lang.Test)
	return nil
}
