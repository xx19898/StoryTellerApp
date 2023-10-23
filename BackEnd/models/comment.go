package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;unique;autoIncrement"`
	TextContent string
	Username    string
	Sender      User  `gorm:"foreignKey:Username;references:Name"`
	Story       Story `gorm:"foreignKey:StoryID"`
	StoryID     uint
}
