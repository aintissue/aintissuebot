package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// Handler for link bot command
func linkCommand(c telebot.Context) error {
	m := c.Message()
	getUserOrCreate(m)

	chat := getChatByTgId(m.Chat.ID)
	lnk := generateLink(chat.Namespace)
	res := fmt.Sprintf(lang.LinkMsg, lnk)

	_, err := bot.Send(m.Chat, res, telebot.NoPreview)
	if err != nil {
		loge(err)
	}

	return nil
}
