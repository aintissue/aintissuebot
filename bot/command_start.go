package bot

import (
	"gopkg.in/telebot.v3"
)

// Handler for start bot command
func startCommand(c telebot.Context) error {
	m := c.Message()

	response := lang.StartMsg

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	newUser(m)

	// if len(m.Payload) > 0 {
	// 	id, err := strconv.Atoi()
	// }

	return nil
}
