package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

// Handler for register bot command
func registerCommand(c telebot.Context) error {
	// m := c.Message()

	// response := lang.StartMsg

	log.Println(c.Chat().Username)

	// _, err := bot.Send(m.Chat, response, telebot.NoPreview)
	// if err != nil {
	// 	loge(err)
	// }

	return nil
}
