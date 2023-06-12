package models

type User struct {
	ID       uint `gorm:"primaryKey;unique;autoIncrement"`
	Name     string
	Password string
	Email    string
}
