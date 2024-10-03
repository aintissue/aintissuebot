package bot

import (
	"gopkg.in/telebot.v3"
)

// Handler for text bot command
func textCommand(c telebot.Context) error {
	m := c.Message()
	var rec *telebot.User
	var msg string
	var chat *Chat
	msgExists := false
	u := getUserOrCreate(m)

	if m.IsReply() {
		_, msgExists = msgs[m.ReplyTo.ID]
	}

	if (m.IsReply() && msgExists) || m.Private() {
		u.MsgCount++
		if err := db.Save(u).Error; err != nil {
			loge(err)
		}
	}

	if m.IsReply() {
		msg = m.Text
		rec = &telebot.User{ID: msgs[m.ReplyTo.ID]}

		_, err := bot.Send(rec, msg, telebot.NoPreview)
		if err != nil {
			loge(err)
		}

		delete(msgs, m.ReplyTo.ID)
	} else if m.Private() {
		chat = getChatById(u.DefaultChatID)
		rec = &telebot.User{ID: chat.OwnerID}
		// msg = fmt.Sprintf("<b><u>%s:</u></b>\n%s",
		// 	m.Sender.FirstName,
		// 	m.Text)

		// mn, err := bot.Send(rec, msg, telebot.NoPreview)
		// if err != nil {
		// 	loge(err)
		// }

		mn, err := bot.Forward(rec, m, telebot.NoPreview)
		if err != nil {
			loge(err)
		}

		msgs[mn.ID] = m.Sender.ID
	}

	return nil
}
