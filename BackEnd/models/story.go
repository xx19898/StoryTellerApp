package models

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;unique;autoIncrement" json:"omitempty"`
	Content  string
	Title    string
	UserID   uint
	Owner    User `gorm:"foreignKey:UserID"`
	Comments []Comment
}
