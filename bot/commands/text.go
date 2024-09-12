package commands

import (
	"log"

	"gopkg.in/telebot.v3"
)

func TextCommand(c telebot.Context) error {
	log.Println("text command")
	return nil
}
