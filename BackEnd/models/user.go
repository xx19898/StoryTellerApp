package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;unique;autoIncrement"`
	Name     string
	Password string
	Email    string
	Roles    []Role `gorm:"many2many:user_roles;"`
}
