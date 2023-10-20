package models

import "gorm.io/gorm"

//TODO: figure out the error here, ID no longer exists in User, switch it to
//TODO: the username, figure out what type of relation this is
type Story struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;unique;autoIncrement" json:"omitempty"`
	Content  string
	Title    string
	Username string
	Owner    User `gorm:"foreignKey:Username;references:Name"`
	Comments []Comment
}
