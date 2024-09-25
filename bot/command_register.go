package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	m := c.Message()
	response := lang.RegisterError
	getUserOrCreate(m)

	if len(m.Payload) > 0 {
		chat := newChat(m.Payload, c.Chat().ID)
		if chat.ID != 0 {
			response = fmt.Sprintf(lang.RegisterDone, generateLink(m.Payload))
		} else {
			response = lang.RegDuplicate
		}
	} else if len(c.Chat().Username) > 0 {
		chat := newChat(c.Chat().Username, c.Chat().ID)
		if chat.ID != 0 {
			response = fmt.Sprintf(lang.RegisterDone, generateLink(c.Chat().Username))
		} else {
			response = lang.RegDuplicate
		}
	}

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
