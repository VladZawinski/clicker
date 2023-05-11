package models

import "gorm.io/gorm"

type User struct {
	Name       string
	Phone      string
	Password   string
	Role       string
	UserClicks []UserClicks `gorm:"foreignKey:UserID"`
	gorm.Model
}
