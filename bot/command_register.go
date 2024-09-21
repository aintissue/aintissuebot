package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	m := c.Message()
	response := "You either have to define your project name with <code>/register ProjectName</code> or you have to have public username / link on your Telegram account or group."

	if len(m.Payload) > 0 {

		newChat(m.Payload, c.Chat().ID)

		response = fmt.Sprintf(`registered

%s`,
			generateLink(m.Payload))

	} else if len(c.Chat().Username) > 0 {

		newChat(c.Chat().Username, c.Chat().ID)

		response = fmt.Sprintf(`registered

%s`,
			generateLink(c.Chat().Username))

	}

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
