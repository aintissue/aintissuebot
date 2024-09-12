package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for text bot command
func textCommand(c telebot.Context) error {
	log.Println(c.Message().Private())
	return nil
}
