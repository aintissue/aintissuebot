package bot

import (
	"gopkg.in/telebot.v3"
)

// Handler for start bot command
func startCommand(c telebot.Context) error {
	m := c.Message()

	response := lang.StartMsg

	u := getUserOrCreate(m)

	if len(u.RefCode) == 0 && len(u.TelUsername) > 0 {
		u.RefCode = u.TelUsername
	} else if len(u.RefCode) == 0 {
		u.RefCode = randomString(10)
	}

	if len(m.Payload) > 0 {
		if m.Payload == "login" {
			response = getLoginLink(u)
		} else {
			c := getChat(m.Payload)
			u.DefaultChatID = c.ID
		}
	} else {
		u.DefaultChatID = 1
	}

	if err := db.Save(u).Error; err != nil {
		loge(err)
	}

	_, err := bot.Send(m.Chat, response, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
