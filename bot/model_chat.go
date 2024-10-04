package bot

import (
	"strings"

	"gorm.io/gorm"
)

// Struct representing Chat object
type Chat struct {
	gorm.Model
	Namespace string `gorm:"size:255;uniqueIndex"`
	OwnerID   uint
	ChatID    int64
}

// Creates new Chat object
func newChat(ns string, owner *User, chatId int64) (*Chat, string) {
	ns = normalizeNs(ns)
	r := ""

	c := &Chat{
		Namespace: ns,
		OwnerID:   owner.ID,
		ChatID:    chatId,
	}

	if owner.Plan == PlanFree && len(owner.getChats()) > 0 {
		return c, lang.MustUpgrade
	}

	if err := db.Save(c).Error; err != nil {
		loge(err)
		if strings.Contains(err.Error(), "UNIQUE") {
			r = lang.RegDuplicate
		}
	}

	return c, r
}

// Fetches Chat object by namespace
func getChat(ns string) *Chat {
	c := &Chat{}
	ns = normalizeNs(ns)

	if err := db.First(c, &Chat{Namespace: ns}).Error; err != nil {
		loge(err)
	}

	return c
}

// Fetches Chat object by ID
func getChatById(id uint) *Chat {
	c := &Chat{}

	if err := db.First(c, id).Error; err != nil {
		loge(err)
	}

	return c
}

// Fetches Chat object by Telegram ID
func getChatByTgId(id int64) *Chat {
	c := &Chat{}

	if err := db.First(c, &Chat{ChatID: id}).Error; err != nil {
		loge(err)
	}

	return c
}

func normalizeNs(ns string) string {
	ns = strings.ToLower(ns)
	ns = strings.ReplaceAll(ns, " ", "")

	return ns
}
