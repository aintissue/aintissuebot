package bot

import "gorm.io/gorm"

// Struct representing Chat object
type Chat struct {
	gorm.Model
	Namespace string `gorm:"size:255;uniqueIndex"`
	OwnerID   int64
}

// Creates new Chat object
func newChat(ns string, ownerId int64) *Chat {
	c := &Chat{
		Namespace: ns,
		OwnerID:   ownerId,
	}

	if err := db.Save(c).Error; err != nil {
		loge(err)
	}

	return c
}

// Fetches Chat object by namespace
func getChat(ns string) *Chat {
	c := &Chat{}

	if err := db.First(c, &Chat{Namespace: ns}).Error; err != nil {
		loge(err)
	}

	return c
}

// Fetches Chat object by ID
func getChatId(id uint) *Chat {
	c := &Chat{}

	if err := db.First(c, id).Error; err != nil {
		loge(err)
	}

	return c
}

// Fetches Chat object by Telegram ID
func getChatByTgId(id int64) *Chat {
	c := &Chat{}

	if err := db.First(c, &Chat{OwnerID: id}).Error; err != nil {
		loge(err)
	}

	return c
}
