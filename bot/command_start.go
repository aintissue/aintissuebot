package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for start bot command
func startCommand(c telebot.Context) error {
	m := c.Message()

	response := lang.StartMsg

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		log.Println(err)
	}

	return nil
}
