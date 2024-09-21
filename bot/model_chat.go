package bot

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Namespace string `gorm:"size:255;uniqueIndex"`
	OwnerID   uint
}

func newChat() *Chat {
	c := &Chat{}

	return c
}
