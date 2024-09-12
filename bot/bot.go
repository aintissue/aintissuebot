package bot

import (
	"log"

	"gopkg.in/telebot.v3"
)

var conf *Config

var bot *telebot.Bot

var lang *Language

// Package init function
func init() {
	conf = initConfig()

	lang = initLang()

	bot = initTelegramBot()
}

// Prepares the environment and runs the bot
func Run() {
	log.Println("Bot started successfully.")

	bot.Start()
}
