package bot

import (
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for text bot command
func textCommand(c telebot.Context) error {
	m := c.Message()
	var rec *telebot.User
	var msg string

	if m.Private() {
		if m.Sender.ID != conf.OwnerId {
			rec = &telebot.User{ID: conf.OwnerId}
			msg = fmt.Sprintf("<b><u>%s:</u></b>\n%s",
				m.Sender.FirstName,
				m.Text)
		} else if m.IsReply() {
			msg = m.Text
			rec = &telebot.User{ID: msgs[m.ReplyTo.ID]}
		}

		mn, err := bot.Send(rec, msg, telebot.NoPreview)
		if err != nil {
			log.Println(err)
		}

		msgs[mn.ID] = m.Sender.ID
	}

	return nil
}
