package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	m := c.Message()

	// response := lang.StartMsg

	log.Println(c.Chat().Username)
	log.Println(m.Payload)

	// _, err := bot.Send(m.Chat, response, telebot.NoPreview)
	// if err != nil {
	// 	loge(err)
	// }

	if len(m.Payload) > 0 {

	} else if len(c.Chat().Username) > 0 {

	} else {

	}

	return nil
}
