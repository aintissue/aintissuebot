package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	m := c.Message()
	response := lang.RegisterError

	if len(m.Payload) > 0 {
		newChat(m.Payload, c.Chat().ID)
		response = fmt.Sprintf(lang.RegisterDone, generateLink(m.Payload))
	} else if len(c.Chat().Username) > 0 {
		newChat(c.Chat().Username, c.Chat().ID)
		response = fmt.Sprintf(lang.RegisterDone, generateLink(c.Chat().Username))
	}

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
