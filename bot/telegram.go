package bot

import (
	"aintissuebot/bot/commands"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

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

	return b
}

func initTelegramCommands(b *telebot.Bot) {
	b.Handle("/start", commands.StartCommand)
	b.Handle(telebot.OnText, commands.TextCommand)
}
