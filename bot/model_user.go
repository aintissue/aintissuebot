package bot

import (
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

// Struct representing User object
type User struct {
	gorm.Model
	TelegramId    int64  `gorm:"uniqueIndex"`
	TelUsername   string `gorm:"size:255"`
	TelFirstName  string `gorm:"size:255"`
	TelLastName   string `gorm:"size:255"`
	DefaultChatID uint
	DefaultChat   *Chat `gorm:"default:1"`
	MsgCount      uint64
}

// Fetches User object by Telegram ID
func getUser(tid int64) *User {
	u := &User{}
	// log.Println(prettyPrint(u))
	if err := db.First(u, &User{TelegramId: tid}).Error; err != nil {
		loge(err)
	}

	return u
}

// func userExists(tid int64) bool {
// 	u := getUser(tid)

// 	return u.ID != 0
// }

// Creates new User object in the database using telebot.Message object
func newUser(m *telebot.Message) *User {
	u := &User{
		TelegramId:   m.Sender.ID,
		TelUsername:  m.Sender.Username,
		TelFirstName: m.Sender.FirstName,
		TelLastName:  m.Sender.LastName,
	}

	if err := db.Save(u).Error; err != nil {
		loge(err)
	}

	return u
}

// Fetches User object and creates it if it doesn't exist
func getUserOrCreate(m *telebot.Message) *User {
	u := getUser(m.Sender.ID)

	if u.ID == 0 {
		u = newUser(m)
	}

	return u
}
