package bot

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open("data/bot.db"), &gorm.Config{})

	if err != nil {
		loge(err)
	}

	if err := db.AutoMigrate(&Chat{}); err != nil {
		panic(err.Error())
	}

	return db
}
