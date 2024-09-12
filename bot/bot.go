package bot

import (
	"gopkg.in/telebot.v3"
)

var conf *Config

var bot *telebot.Bot

// Package init function
func init() {
	conf = initConfig()

	bot = initTelegramBot()
}

// Prepares the environment and runs the bot
func Run() {
	// log.Println(conf.TelegramKey)
	bot.Start()
}
