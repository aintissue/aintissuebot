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

	cmd := telebot.Command{Text: "help", Description: "Some help command"}

	cmds, err := b.Commands()
	if err != nil {
		log.Println(err)
	}

	log.Println(prettyPrint(cmds))

	sc := telebot.CommandScope{
		Type: telebot.CommandScopeAllPrivateChats,
	}

	err = b.SetCommands([]telebot.Command{cmd}, sc)
	// err = b.SetCommands([]telebot.Command{}, sc)
	if err != nil {
		log.Println(err)
	}

	return b
}

// Initializes Telegram bot commands
func initTelegramCommands(b *telebot.Bot) {
	b.Handle("/start", startCommand)
	b.Handle("/register", registerCommand)
	b.Handle(telebot.OnText, textCommand)
}
