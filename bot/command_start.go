package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for start bot command
func startCommand(c telebot.Context) error {
	m := c.Message()

	log.Println(m.Text)

	response := lang.StartMsg

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
