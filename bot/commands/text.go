package commands

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for text bot command
func TextCommand(c telebot.Context) error {
	log.Println(c.Message().Text)
	return nil
}
