package commands

import (
	"log"

	"gopkg.in/telebot.v3"
)

func StartCommand(c telebot.Context) error {
	log.Println("start command")
	return nil
}
