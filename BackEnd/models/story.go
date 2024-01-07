package models

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;unique;autoIncrement"`
	Content  string
	Title    string
	Username string
	Owner    User `gorm:"foreignKey:Username;references:Name"`
	Comments []Comment
}
