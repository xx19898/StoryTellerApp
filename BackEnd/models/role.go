package models

type Role struct {
	Name string
	ID   int `gorm:"primaryKey;unique;autoIncrement"`
}
