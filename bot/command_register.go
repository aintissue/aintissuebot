package bot

import (
	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	m := c.Message()

	if len(m.Payload) > 0 {
		newChat(m.Payload, c.Chat().ID)
	} else if len(c.Chat().Username) > 0 {
		newChat(c.Chat().Username, c.Chat().ID)
	} else {

	}

	response := "registered"
	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
