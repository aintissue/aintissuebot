package bot

import (
	"log"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

var conf *Config

var bot *telebot.Bot

var lang *Language

var msgs map[int]int64

var db *gorm.DB

// Package init function
func init() {
	conf = initConfig()

	lang = initLang()

	bot = initTelegramBot()

	msgs = make(map[int]int64)

	db = initDb()
}

// Prepares the environment and runs the bot
func Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logs("Bot started successfully. 🚀")

	bot.Start()
}
