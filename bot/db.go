package bot

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initializes DB object
func initDb() *gorm.DB {
	var db *gorm.DB
	var err error

	if conf.Dev {
		db, err = gorm.Open(sqlite.Open(DbPath), &gorm.Config{})
	} else {
	}

	if err != nil {
		loge(err)
	}

	if err := db.AutoMigrate(&Chat{}, &User{}); err != nil {
		panic(err.Error())
	}

	return db
}
