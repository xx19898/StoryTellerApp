package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;unique;autoIncrement" json:"omitempty"`
	TextContent string
	UserID      uint
	Sender      User `gorm:"foreignKey:UserID"`
	Story       Story
	StoryID     uint
}
