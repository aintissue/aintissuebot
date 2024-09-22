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

	u := getUserOrCreate(m)

	if len(m.Payload) > 0 {
		c := getChat(m.Payload)
		u.DefaultChatID = c.ID
	} else {
		u.DefaultChatID = 1
	}

	if err := db.Save(u).Error; err != nil {
		loge(err)
	}

	return nil
}
