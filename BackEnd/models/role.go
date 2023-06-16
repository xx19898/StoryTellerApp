package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
	ID   int `gorm:"primaryKey;unique;autoIncrement"`
}
