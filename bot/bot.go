package bot

import (
	"gopkg.in/telebot.v3"
)

var conf *Config

var bot *telebot.Bot

var lang *Language

var msgs map[int]int64

// Package init function
func init() {
	conf = initConfig()

	lang = initLang()

	bot = initTelegramBot()

	msgs = make(map[int]int64)
}

// Prepares the environment and runs the bot
func Run() {
	logs("Bot started successfully. 🚀")

	bot.Start()
}
