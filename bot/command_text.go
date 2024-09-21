package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// Handler for text bot command
func textCommand(c telebot.Context) error {
	m := c.Message()
	var rec *telebot.User
	var msg string

	if m.IsReply() {
		msg = m.Text
		rec = &telebot.User{ID: msgs[m.ReplyTo.ID]}

		_, err := bot.Send(rec, msg, telebot.NoPreview)
		if err != nil {
			loge(err)
		}

		delete(msgs, m.ReplyTo.ID)
	} else if m.Private() {
		u := getUser(m.Chat.ID)
		c := getChatId(u.DefaultChatID)
		rec = &telebot.User{ID: c.OwnerID}
		msg = fmt.Sprintf("<b><u>%s:</u></b>\n%s",
			m.Sender.FirstName,
			m.Text)

		mn, err := bot.Send(rec, msg, telebot.NoPreview)
		if err != nil {
			loge(err)
		}

		msgs[mn.ID] = m.Sender.ID
	}

	return nil
}
