package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;unique;autoIncrement"`
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}
