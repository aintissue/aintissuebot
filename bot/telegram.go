package bot

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

// Initializes Telegram bot
func initTelegramBot() *telebot.Bot {
	b, err := telebot.NewBot(telebot.Settings{
		Token:     conf.TelegramKey,
		Poller:    &telebot.LongPoller{Timeout: 30 * time.Second},
		Verbose:   false,
		ParseMode: "html",
	})

	if err != nil {
		log.Fatal(err)
	}

	initTelegramCommands(b)

	// cmd := telebot.Command{Text: "register", Description: "Register your account or group"}
	// cmd1 := telebot.Command{Text: "link", Description: "Show your customer support link"}

	// cmds, err := b.Commands()
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(prettyPrint(cmds))

	sc := telebot.CommandScope{
		Type: telebot.CommandScopeAllPrivateChats,
	}

	// err = b.SetCommands([]telebot.Command{cmd, cmd1}, sc)
	err = b.SetCommands([]telebot.Command{}, sc)
	if err != nil {
		log.Println(err)
	}

	return b
}

// Initializes Telegram bot commands
func initTelegramCommands(b *telebot.Bot) {
	b.Handle("/start", startCommand)
	b.Handle("/register", registerCommand)
	b.Handle("/link", linkCommand)
	b.Handle(telebot.OnText, textCommand)
	b.Handle(telebot.OnPhoto, textCommand)
	b.Handle(telebot.OnMedia, textCommand)
}
