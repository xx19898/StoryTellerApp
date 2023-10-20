package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"primaryKey;unique;" json:"username,omitempty" binding:"required" `
	Password string `json:"password,omitempty" binding:"required"`
	Email    string `json:"email,omitempty"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"omitempty"`
}
